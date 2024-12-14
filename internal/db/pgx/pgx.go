package pgx

import (
	"context"
	"net/url"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBConfig struct {
	Host         string
	Port         string
	Name         string
	User         string
	Password     string
	DisableTLS   bool
	MaxOpenConns int
	MaxIdleConns int
}

// openDB opens a database connection.
func OpenDBConnection(ctx context.Context, cfg *DBConfig) (*pgx.Conn, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	DSN := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host + ":" + cfg.Port,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	config, err := pgx.ParseConfig(DSN.String())
	if err != nil {
		return nil, err
	}

	db, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return db, nil
}
