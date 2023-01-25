package main

import (
	"log"

	"github.com/IDL13/echo/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
