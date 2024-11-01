package main

import (
	"context"
	"httpserver/internal/api"
	"httpserver/internal/config"
	"log"
	"os"
)

type Config struct {
	Host string
	Port string
}

// this vars are set by the compiler via ldflags. e.g. go build -ldflags "-X main.build=production -X main.version=v1.0.0"
var build = "development"
var version = "v0.0.0"

func main() {
	// Create a context for the application.
	ctx := context.Background()

	// Create a logger for the application.
	logger := log.New(log.Writer(), log.Prefix(), log.Flags()) // adjust the logger to your needs

	// Load the configuration for the application.
	cfg := config.LoadConfig() // any config for now

	// Call the run function to start the application.
	// Start the application.
	logger.Println("Starting the application...", "build", build, "version", version)

	if err := api.Run(ctx, logger, cfg); err != nil {
		logger.Println(err)
		os.Exit(1)
	}

}
