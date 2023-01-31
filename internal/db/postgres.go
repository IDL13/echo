package db

import (
	"context"
	"fmt"
	"log"

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

	q := `INSERT INTO card (number, date, cvv) VALUES ($1, $2, $3) RETURNING number`
	err = conn.QueryRow(context.TODO(), q, date.Number, date.Date, date.CVV).Scan(&date.Number, &date.Date, &date.CVV)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close(context.TODO())

	fmt.Println("successful ading")
	return nil
}

func (r *Repository) FindAll() (err error) {
	cfg := config.GetConf()
	conn, err := r.client.NewClient(*cfg)
	if err != nil {
		log.Println("error")
	}

	q := `SELECT number FROM public.card`
	all, _ := conn.Query(context.Background(), q)
	if err != nil {
		log.Println("error")
	}

	for all.Next() {
		var card Card

		err = all.Scan(&card.Number)
		if err != nil {
			fmt.Println("error in FindAll")
		}
		log.Println(card.Number)
	}
	return nil
}
