package main

import (
	"log"

	"github.com/boldd/internal/api"
	"github.com/boldd/internal/config"
)

func main() {
	// Load the configuration
	cfg, err := config.Load("$HOME/.config/boldd")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server := api.NewApplication(cfg)
	go server.Shutdown()

	if err := server.Run(); err != nil {
		log.Fatalf("🔴 Failed to start server")
	}

	<-server.Done
	log.Println("👋 Server shutdown gracefully...")
}
