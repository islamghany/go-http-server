package api

import (
	"context"
	"fmt"
	"httpserver/internal/config"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// The run function is like the main function, except that it takes in operating system fundamentals as arguments, and returns an error.
func Run(
	ctx context.Context,
	logger *log.Logger,
	cfg *config.Config,
) error {

	// Starting the DB
	// db, err := db.NewDB(cfg)

	// Starting the debug server if needed in a go routine, and handlign the wait group

	// Starting the HTTP server with graceful shutdown
	srv := NewServer(logger, nil)

	server := &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: srv,
	}
	shutdownError := make(chan error)

	go func() {
		logger.Println("Starting the HTTP server on", server.Addr)
		shutdownError <- server.ListenAndServe()
	}()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-shutdownError:
		return fmt.Errorf("error starting the server: %w", err)
	case err := <-shutdownChan:
		{
			logger.Println("Shutting down the server...", err)
			// create a timeout context that carries the deadline
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				server.Close()
				return fmt.Errorf("error shutting down the server gracefully: %w", err)
			}
		}
	}

	return nil

}
