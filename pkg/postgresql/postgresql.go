package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func New() *Client {
	return &Client{}
}

type Client struct {
}

func (c *Client) NewClient() (conn *pgx.Conn, err error) {
	conn, err = pgx.Connect(context.Background(), "postgres://echo:root6895963@localhost:5432/echo")
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
