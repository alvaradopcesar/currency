package postgres

import (
	"context"
	"database/sql"
	"sync"
	"time"
)

var (
	postgresClient *Client
	postgresOnce   sync.Once
)

// Client is a client for the PostgreSQL db engine.
type Client struct {
	*sql.DB
}

// NewPostgresClient returns a new client for postgres.
func NewPostgresClient(source string) *Client {
	postgresOnce.Do(func() {
		db, err := sql.Open("postgres", source)
		if err != nil {
			panic(err)
		}

		err = db.PingContext(context.Background())
		if err != nil {
			panic(err)
		}

		// Limit the number of open connections to avoid
		// memory problems in the database.
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(2 * time.Minute)

		postgresClient = &Client{db}
	})

	return postgresClient
}

// ViewStats returns the status of the db.
func (c *Client) ViewStats() sql.DBStats {
	return c.Stats()
}
