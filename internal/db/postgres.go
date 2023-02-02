package db

import (
	"context"
	"fmt"
	"log"

	config "github.com/IDL13/echo/internal/config"
	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
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
		fmt.Println("Error from NewClient")
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

func (r *Repository) FindAll() (mas []string) {
	cfg := config.GetConf()
	conn, err := r.client.NewClient(*cfg)
	if err != nil {
		log.Println("error")
	}

	var m []string

	q := `SELECT number FROM public.card`
	all, _ := conn.Query(context.Background(), q)
	if err != nil {
		log.Println("error")
	}

	for all.Next() {
		var card Card

		err = all.Scan(&card.Number)
		m = append(m, card.Number)
		if err != nil {
			fmt.Println("error in FindAll")
		}
		log.Println(card.Number)
	}
	return m
}

func (r *Repository) FindOne(n *unmarshal.Name) (err error) {
	cfg := config.GetConf()
	conn, err := r.client.NewClient(*cfg)
	if err != nil {
		log.Println("Error from configfile")
	}

	var card Card

	q := `SELECT * FROM public.card WHERE number = $1`
	err = conn.QueryRow(context.Background(), q, n.Name).Scan(&card.Number, &card.Date, &card.CVV)
	if err != nil {
		log.Println("Error in QueryRow")
	}
	log.Printf("Number:%s", card.Number)
	log.Printf("Date:%s", card.Number)
	log.Printf("CVV:%s", card.Number)

	return nil

}
