package db

import (
	"context"
	"fmt"

	"github.com/IDL13/echo/pkg/postgresql"
)

func New() *Repository {
	return &Repository{}
}

type Repository struct {
	client postgresql.Client
}

func (r *Repository) Insert() error {
	conn, err := r.client.NewClient()
	if err != nil {
		fmt.Println("E-R-R-O-R")
	}

	num1 := "jkkjhjk"
	num2 := "sdfgdfg"
	num3 := "agsdgsg"

	err = conn.QueryRow(context.Background(), `insert into "Card" ("Number", "Date", "CVV") values ($1, $2, $3)`, num1, num2, num3).Scan(&num1, &num2, &num3)
	if err != nil {
		fmt.Println("QueryRow failed")
	}

	defer conn.Close(context.Background())

	fmt.Println("successful ading")
	return nil

	// n1 := "1"
	// n2 := "2"
	// n3 := "3"
	// q := `insert into "Card" ("Number", "Date", "CVV") values($1, $2, $3)`
	// fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	// err := r.client.QueryRow(context.Background(), q, n1, n2, n3).Scan(&n1, &n2, &n3)
	// if err != nil {
	// 	fmt.Printf("Create error %d", err)
	// }
	// fmt.Println("Secessfull insert")
	// return nil
}

// func (r *Repository) FindAll(ctx context.Context) (c []Card, err error) {
// 	q := `SELECT Number, Date, CVV FROM public.Card;`

// 	rows, err := r.client.Query(ctx, q)
// 	if err != nil {
// 		return nil, err
// 	}

// 	cards := make([]Card, 0)

// 	for rows.Next() {
// 		var car Card

// 		err = rows.Scan(&car.Number, &car.Date, &car.CVV)
// 		if err != nil {
// 			return nil, err
// 		}

// 		cards = append(cards, car)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return cards, nil
// }
