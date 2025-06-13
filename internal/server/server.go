package server

type Application struct {
	done chan bool
}

func NewApplication() *Application {
	return &Application{
		done: make(chan bool),
	}
}

func (app *Application) Start() error {
	// Initialize and start the application
	// This is where you would set up routes, middleware, etc.
	return nil
}

func (app *Application) Shutdown() {
	// shutdown the application gracefully
}

func (app *Application) Run() error {
	// Start the application and block until it is stopped
	if err := app.Start(); err != nil {
		return err
	}
	return nil
}
