package db

import (
	"context"
	"errors"
	"fmt"

	// "fmt"
	"log"

	"github.com/IDL13/echo/internal/config"
	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
	"github.com/IDL13/echo/pkg/postgresql"
	"github.com/IDL13/echo/pkg/utils"
	"github.com/jackc/pgx/v5/pgconn"
)

func New() Repository {
	r := repository{}
	return &r
}

type repository struct {
	client postgresql.Client
}

func (r *repository) Insert(ctx context.Context, card *encryption.Date) error {
	cfg := config.GetConf()
	conn, err := postgresql.NewClient(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	q := `INSERT INTO card (number, date, cvv) VALUES ($1, $2, $3) RETURNING number`
	err = conn.QueryRow(context.TODO(), q, card.Number, card.Date, card.CVV).Scan(&card.Number, &card.Date, &card.CVV)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
		}

		utils.Loger(err)
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) (mas []Card, err error) {
	cfg := config.GetConf()
	conn, err := postgresql.NewClient(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	var m []Card
	// m := make([]Card, 0)

	q := `SELECT number, date, cvv FROM card`
	all, _ := conn.Query(ctx, q)
	if err != nil {
		utils.Loger(err)
	}

	for all.Next() {
		var card Card

		err = all.Scan(&card.Number, &card.Date, &card.CVV)
		if err != nil {
			utils.Loger(err)
		}
		log.Println(card.Number)

		m = append(m, card)
	}
	return m, nil
}

func (r *repository) FindOne(ctx context.Context, number *unmarshal.Name) error {
	cfg := config.GetConf()
	conn, err := postgresql.NewClient(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	var card Card

	q := `SELECT * FROM public.card WHERE number = $1`

	err = conn.QueryRow(ctx, q, number).Scan(&card.Number, &card.Date, &card.CVV)
	if err != nil {
		utils.Loger(err)
	}
	fmt.Printf("Number:%s", card.Number)
	fmt.Printf("Date:%s", card.Number)
	fmt.Printf("CVV:%s", card.Number)

	return nil
}
