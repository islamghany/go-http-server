package api

import (
	"context"
	"fmt"
	"httpserver/internal/config"
	pgx "httpserver/internal/db/pgx"
	db "httpserver/internal/db/sqlc"
	"httpserver/internal/logger"
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
	logger *logger.Logger,
	cfg *config.Config,
) error {

	// Starting the DB
	dbConn, err := pgx.OpenDBConnection(ctx, &pgx.DBConfig{
		Host:         cfg.DBHost,
		Port:         cfg.DBPort,
		Name:         cfg.DBName,
		User:         cfg.DBUser,
		Password:     cfg.DBPassword,
		DisableTLS:   cfg.DBDisableTLS,
		MaxOpenConns: cfg.DBMaxOpenConns,
	})
	if err != nil {
		return fmt.Errorf("error opening the database connection: %w", err)
	}
	defer dbConn.Close(ctx)
	store := db.New(dbConn)
	// Starting the debug server if needed in a go routine, and handlign the wait group

	// Starting the HTTP server with graceful shutdown
	srv := NewServer(logger, cfg, store)

	server := &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: srv,
	}
	shutdownError := make(chan error)

	go func() {
		// logger.Println("Starting the HTTP server on", server.Addr)
		logger.Info(ctx, "Starting the HTTP server on", "address", server.Addr)
		shutdownError <- server.ListenAndServe()
	}()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-shutdownError:
		return fmt.Errorf("error starting the server: %w", err)
	case err := <-shutdownChan:
		{
			logger.Info(ctx, "Shutting down the server...", "signal", err)
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
