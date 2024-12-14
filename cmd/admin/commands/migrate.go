package commands

import (
	"context"
	"fmt"
	"httpserver/internal/config"
	"httpserver/pkg/logger"
	"net/url"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDBTo(log *logger.Logger, cfg *config.Config, down bool) error {
	ctx := context.Background()

	q := make(url.Values)
	ssl := "require"
	if cfg.DBDisableTLS {
		ssl = "disable"
	}
	q.Set("sslmode", ssl)
	q.Set("timezone", "utc")
	DSN := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:     cfg.DBHost + ":" + cfg.DBPort,
		Path:     cfg.DBName,
		RawQuery: q.Encode(),
	}
	m, err := migrate.New(
		"file://internal/db/migrations",
		DSN.String(),
	)

	if err != nil {
		return err
	}

	if down {
		err := m.Steps(-1)
		if err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("m.Down: %w", err)
		}
		v, _, er := m.Version()

		if er != nil && er != migrate.ErrNilVersion {
			return fmt.Errorf("m.Version: %w", er)
		}

		log.Info(ctx, "Migrationg down", "version", v)

	} else {
		err := m.Up()
		if err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("m.Up: %w", err)
		}
		v, _, er := m.Version()
		if er != nil {
			return fmt.Errorf("m.Version: %w", er)
		}
		log.Info(ctx, "Migrationg up", "version", v)
	}

	return nil
}
