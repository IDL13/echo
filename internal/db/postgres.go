package db

import (
	"context"
	"fmt"

	// "fmt"
	"log"

	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
	"github.com/IDL13/echo/pkg/postgresql"
	"github.com/IDL13/echo/pkg/utils"
)

// func New() *Repository {
// 	return &Repository{}
// }

type repository struct {
	client postgresql.Client
}

// type Repository struct {
// 	client postgresql.Client
// }

func (r *repository) Insert(ctx context.Context, card *encryption.Date) error {
	q := `INSERT INTO card (number, date, cvv) VALUES ($1, $2, $3) RETURNING number`
	if err := r.client.QueryRow(ctx, q, card.Number, card.Date, card.CVV).Scan(&card.Number, &card.Date, &card.CVV); err != nil {
		return err
	}

	return nil
}

// func (r *Repository) Insert(date *encryption.Date) error {
// 	cfg := config.GetConf()
// 	conn, err := r.client.NewClient(*cfg)
// 	if err != nil {
// 		utils.Loger(err)
// 	}

// 	q := `INSERT INTO card (number, date, cvv) VALUES ($1, $2, $3) RETURNING number`
// 	err = conn.QueryRow(context.TODO(), q, date.Number, date.Date, date.CVV).Scan(&date.Number, &date.Date, &date.CVV)
// 	if err != nil {
// 		utils.Loger(err)
// 	}

// 	defer conn.Close(context.TODO())

// 	log.Println("successful ading")
// 	return nil
// }

func (r *repository) FindAll(ctx context.Context) (mas []Card, err error) {
	var m []Card
	// m := make([]Card, 0)

	q := `SELECT number, date FROM public.card`
	all, _ := r.client.Query(ctx, q)
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

// func (r *Repository) FindAll() (mas []string) {
// 	cfg := config.GetConf()
// 	conn, err := r.client.NewClient(*cfg)
// 	if err != nil {
// 		utils.Loger(err)
// 	}

// 	var m []string

// 	q := `SELECT number FROM public.card`
// 	all, _ := conn.Query(context.Background(), q)
// 	if err != nil {
// 		utils.Loger(err)
// 	}

// 	for all.Next() {
// 		var card Card

// 		err = all.Scan(&card.Number)
// 		m = append(m, card.Number)
// 		if err != nil {
// 			utils.Loger(err)
// 		}
// 		log.Println(card.Number)
// 	}
// 	return m
// }

func (r *repository) FindOne(ctx context.Context, number *unmarshal.Name) error {
	var card Card

	q := `SELECT * FROM public.card WHERE number = $1`

	err := r.client.QueryRow(ctx, q, number).Scan(&card.Number, &card.Date, &card.CVV)
	if err != nil {
		utils.Loger(err)
	}
	fmt.Printf("Number:%s", card.Number)
	fmt.Printf("Date:%s", card.Number)
	fmt.Printf("CVV:%s", card.Number)

	return nil
}

// func (r *Repository) FindOne(n *unmarshal.Name) (err error) {
// 	cfg := config.GetConf()
// 	conn, err := r.client.NewClient(*cfg)
// 	if err != nil {
// 		utils.Loger(err)
// 	}

// 	var card Card

// 	q := `SELECT * FROM public.card WHERE number = $1`
// 	err = conn.QueryRow(context.Background(), q, n.Name).Scan(&card.Number, &card.Date, &card.CVV)
// 	if err != nil {
// 		utils.Loger(err)
// 	}
// 	fmt.Printf("Number:%s", card.Number)
// 	fmt.Printf("Date:%s", card.Number)
// 	fmt.Printf("CVV:%s", card.Number)

// 	return nil
// }
