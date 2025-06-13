package main

import (
	"log"

	"github.com/boldd/internal/config"
	"github.com/boldd/internal/server"
)

func main() {
	// Load the configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server := server.NewApplication(cfg)
	go server.Shutdown()

	if err := server.Run(); err != nil {
		log.Fatalf("🔴 Failed to start server")
	}

	<-server.Done
	log.Println("👋 Server shutdown gracefully...")
}
