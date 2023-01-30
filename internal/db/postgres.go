package db

import (
	"context"
	"fmt"

	config "github.com/IDL13/echo/internal/config"
	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/pkg/postgresql"
)

func New() *Repository {
	return &Repository{}
}

type Repository struct {
	client postgresql.Client
}

func (r *Repository) Insert(date *encryption.Date) error {
	cfg := config.GetConf()
	conn, err := r.client.NewClient(*cfg)
	if err != nil {
		fmt.Println("E-R-R-O-R")
	}

	err = conn.QueryRow(context.TODO(), `INSERT INTO card (number, date, cvv) VALUES ($1, $2, $3) RETURNING number`, date.Number, date.Date, date.CVV).Scan(&date.Number, &date.Date, &date.CVV)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close(context.TODO())

	fmt.Println("successful ading")
	return nil
}
