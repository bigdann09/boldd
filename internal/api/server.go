package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/boldd/internal/api/middlewares"
	"github.com/boldd/internal/api/routes"
	"github.com/boldd/internal/api/services"
	"github.com/boldd/internal/config"
	"github.com/boldd/internal/infrastructure/monitoring"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Application struct {
	Done     chan bool
	engine   *gin.Engine
	server   *http.Server
	config   *config.Config
	services *services.Service
	metrics  *monitoring.Metrics
}

func NewApplication(cfg *config.Config) *Application {
	// switch to release mode for production env
	if cfg.ApplicationConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := &Application{
		Done:   make(chan bool, 1),
		engine: gin.Default(),
		config: cfg,
	}

	// register services
	app.services = services.NewServices(app.config)
	app.metrics = monitoring.NewMetrics()

	// set up the server
	app.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ApplicationConfig.Port),
		Handler:      app.registerroutes(),
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return app
}

func (app *Application) Shutdown() {
	// shutdown the application gracefully
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shut down services running
	if err := app.server.Shutdown(ctx); err != nil {
		app.services.Logger.Fatal("Server shutdown error", zap.Error(err))
		return
	}

	// shutdown database service
	sqlDB, err := app.services.DB.DB()
	if err != nil {
		app.services.Logger.Fatal("Could not get database instance", zap.Error(err))
	}

	// close database
	if err := sqlDB.Close(); err != nil {
		app.services.Logger.Fatal("could not shutdown database service", zap.Error(err))
	}

	// shutdown redis
	if err := app.services.Redis.Conn().Close(); err != nil {
		app.services.Logger.Fatal("could not shutdown redis service", zap.Error(err))
	}

	<-ctx.Done()
	app.services.Logger.Info("Shutting down server...")

	app.Done <- true
}

func (app *Application) Run() error {
	// Start the application and block until it is stopped
	app.services.Logger.Info("✅ Server started successfully")
	if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (app *Application) registerroutes() *gin.Engine {
	// register middlewares
	middleware := middlewares.NewMiddleware(app.engine, app.services)
	middleware.Register(app.config, app.metrics)

	// register routes
	routes := routes.NewRouter(app.engine, app.services, middleware)
	engine := routes.SetupRoutes()
	return engine
}
