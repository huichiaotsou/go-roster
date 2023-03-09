package main

import (
	"github.com/huichiaotsou/go-roster/cmd/server"
	"github.com/huichiaotsou/go-roster/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load config
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
		return
	}

	// Start the server
	server.NewServer().Start()
}
