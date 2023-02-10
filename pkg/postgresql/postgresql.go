package postgresql

import (
	"context"
	"fmt"
	"os"

	config "github.com/IDL13/echo/internal/config"
	"github.com/jackc/pgx/v5"
)

func New() *Client {
	return &Client{}
}

type Client struct {
}

func (c *Client) NewClient(cfg config.Config) (conn *pgx.Conn, err error) {
	q := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	conn, err = pgx.Connect(context.Background(), q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn, nil
}
