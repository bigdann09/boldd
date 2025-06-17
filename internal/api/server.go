package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/boldd/internal/api/handlers"
	"github.com/boldd/internal/api/routes"
	"github.com/boldd/internal/config"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Done   chan bool
	engine *gin.Engine
	server *http.Server
}

func NewApplication(cfg *config.Config) *Application {
	app := &Application{
		Done:   make(chan bool, 1),
		engine: gin.Default(),
	}

	// set up the server
	// fmt.Sprintf(":%d", cfg.Application.Port)
	app.server = &http.Server{
		Addr:         ":8009",
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
		log.Printf("Server shutdown error: %v\n", err)
		return
	}

	<-ctx.Done()
	log.Println("Shutting down server...")

	app.Done <- true
}

func (app *Application) Run() error {
	// Start the application and block until it is stopped
	log.Printf("✅ Server started successfully")
	if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (app *Application) registerroutes() *gin.Engine {
	routes := routes.NewRouter(app.engine)
	engine := routes.SetupRoutes()
	return engine
}
