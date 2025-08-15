package main

import (
	"log"

	"github.com/boldd/internal/api"
	"github.com/boldd/internal/config"
)

// @contact.name	Daniel Ibok
// @contact.url	https://bigdann.vercel.com
// @contact.email	dann.dev09@gmail.com

// @securityDefinitions.apiKey	BearerAuth
// @scheme						bearer
// @in							header
//
//	@name						Authorization
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

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MiwiRW1haWwiOiJkYW5uQGdtYWlsLmNvbSIsImlzcyI6IjIiLCJzdWIiOiIyIiwiZXhwIjoxNzU1MzAwMjYxLCJpYXQiOjE3NTUyOTMwNjEsImp0aSI6IjIifQ.o5Caioq3eVkheyGQHMS3jPIMznDOSZXAiOTEno89Fs4
