package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("[SERVER STARTED]")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler)

	e.Logger.Fatal(e.Start(":8080"))
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "[SERVER STARTED]")
}
