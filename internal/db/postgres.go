package db

import (
	"context"
	"fmt"

	"github.com/IDL13/echo/pkg/postgresql"
)

type Repository struct {
	client postgresql.Client
}

func (r *Repository) Create(ctx context.Context, c *Card) error {
	var n string
	q := `INSERT INTO Card (Number, Date, CVV) VALUES ($1, $2, $3) RETURNING Number`

	if err := r.client.QueryRow(ctx, q, "zero", "zero", "zero").Scan(&n); err != nil {
		fmt.Printf("Create error %d", err)
	}
	return nil
}

func (r *Repository) FindAll(ctx context.Context) (c []Card, err error) {
	q := `SELECT Number, Date, CVV FROM public.Card;`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	cards := make([]Card, 0)

	for rows.Next() {
		var car Card

		err = rows.Scan(&car.Number, &car.Date, &car.CVV)
		if err != nil {
			return nil, err
		}

		cards = append(cards, car)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}
