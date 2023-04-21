package main

import (
	"github.com/IDL13/echo/internal/app"
	"github.com/IDL13/echo/pkg/utils"
)

func main() {
	a, err := app.New()
	if err != nil {
		utils.Loger(err)
	}

	a.Run()
}
