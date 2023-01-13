package postgresql

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	reconection "github.com/IDL13/echo/pkg/utils"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v5"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	err = reconection.Tries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	return pool, nil
}
