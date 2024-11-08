package main

import (
	"context"
	"fmt"
	"httpserver/api"
	"httpserver/internal/config"
	"httpserver/internal/web"
	"httpserver/pkg/logger"
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

	traceIDFunc := func(ctx context.Context) string {
		return web.GetTracerID(ctx)
	}

	// Create a logger for the application.
	// logFile, err := os.OpenFile("logs.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatalf("Failed to open log file: %v", err)
	// }
	// defer logFile.Close()
	// multiWriter := io.MultiWriter(os.Stdout, logFile)
	minLevel := logger.LevelInfo
	if build == "development" {
		minLevel = logger.LevelDebug
	}
	// callback functions for the logger events
	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			// send the error to the error tracking service
			fmt.Println("Error:", r.Message)
		},
	}

	logger := logger.NewWithEvents(os.Stdout, minLevel, "your-app-name", traceIDFunc, events)

	// Load the configuration for the application.
	cfg := config.LoadConfig() // any config for now

	// Call the run function to start the application.
	// Start the application.
	logger.Info(ctx, "Starting the application...", "build", build, "version", version)

	if err := api.Run(ctx, logger, cfg); err != nil {
		logger.Error(ctx, "Failed to start the application.", "error", err)
		os.Exit(1)
	}

}
