package main

import (
	"log"

	"github.com/boldd/internal/api"
	"github.com/boldd/internal/config"
)

// @contact.name	Daniel Ibok
// @contact.url	https://bigdann.vercel.com
// @contact.email	dann.dev09@gmail.com
func main() {
	// Load the configuration
	path, err := config.LoadConfigPath()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.Load(path)
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
