package main

import (
	"context"
	"database/sql"
	"fmt"
	"httpserver/cmd/admin/commands"
	"httpserver/internal/config"
	"httpserver/pkg/logger"
	"net/url"
	"os"
	"strings"
)

var version string = "v0.0.0"
var build string = "development"

func main() {

	ctx := context.Background()
	traceIDFunc := func(ctx context.Context) string {
		return "traceID"
	}
	log := logger.New(os.Stdout, logger.LevelInfo, "admin", traceIDFunc)

	// Load the configuration for the application.
	cfg := config.LoadConfig()

	log.Info(ctx, "Starting admin...", "build", build, "version", version)

	if err := run(log, cfg, os.Args[1]); err != nil {
		log.Error(ctx, "Failed to start the application.", "error", err)
		os.Exit(1)
	}
}

func run(log *logger.Logger, cfg *config.Config, cmd string) error {
	fmt.Println("Running command:", cmd)
	if cmd == "" {
		return fmt.Errorf("no command provided")
	}

	// initialize the db connection
	// todo
	db, err := Connect(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}
	defer db.Close()
	cmds := strings.Split(cmd, ",")
	for _, c := range cmds {
		var err error
		switch c {
		case "migrate", "migrate-up":
			// latest version
			err = commands.MigrateDBTo(log, cfg, false)
			break
		case "migrate-down":
			err = commands.MigrateDBTo(log, cfg, true)
		case "seed":
			err = commands.SeedDB(cfg, db)
			break
		default:
			return fmt.Errorf("unknown command: %s", c)
		}
		if err != nil {
			log.Error(context.Background(), "Failed to run command", "command", c, "error", err)
		}

	}

	return nil

}

func Connect(cfg *config.Config) (*sql.DB, error) {
	q := make(url.Values)
	ssl := "require"
	if cfg.DBDisableTLS {
		ssl = "disable"
	}
	q.Set("sslmode", ssl)
	q.Set("timezone", "utc")

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?%s",
		cfg.DBUser,
		url.QueryEscape(cfg.DBPassword),
		cfg.DBHost+":"+cfg.DBPort,
		cfg.DBName,
		q.Encode(),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return db, nil
}
