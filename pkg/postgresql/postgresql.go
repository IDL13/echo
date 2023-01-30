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

// type Client interface {
// 	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
// 	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
// 	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
// 	Begin(ctx context.Context) (pgx.Tx, error)
// }

// func NewClient(ctx context.Context, maxAttempts int, con config.Config) (pool *pgxpool.Pool, err error) {
// 	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", con.Username, con.Password, con.Host, con.Port, con.Database)

// 	err = reconection.Tries(func() error {
// 		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
// 		defer cancel()

// 		pool, err = pgxpool.Connect(ctx, dsn)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}, maxAttempts, 5*time.Second)

// 	if err != nil {
// 		log.Fatal("error do with tries postgresql")
// 	}

// 	return pool, nil
// }
