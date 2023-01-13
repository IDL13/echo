package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Database(num string, date string, cvv string) {
	conn, err := pgx.Connect(context.Background(), "postgres://echo:root6895963@localhost:5432/echo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), `insert into "Card" ("Number", "Date", "CVV") values($1, $2, $3)`, num, date, cvv).Scan(&num, &date, &cvv)
	if err != nil {
		fmt.Println("QueryRow failed")
	}

	fmt.Println("successful ading")
}
