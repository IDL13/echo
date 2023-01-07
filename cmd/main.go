package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

type Card struct {
	Number string `json:"number"`
	Date   string `json:"date"`
	CVV    string `json:"CVV"`
}

func main() {

	fmt.Println("[SERVER STARTED]")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", getHandler)
	e.POST("/post", postHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func getHandler(c echo.Context) error {
	err := c.String(http.StatusOK, "[SERVER STARTED]")
	if err != nil {
		return err
	}
	return nil
}

func postHandler(c echo.Context) error {
	var card Card

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &card)

	// Convert Number, Date, CVV
	number := []byte(card.Number)
	date := []byte(card.Date)
	cvv := []byte(card.CVV)

	cost := 10

	numberHash, _ := bcrypt.GenerateFromPassword(number, cost)
	dateHash, _ := bcrypt.GenerateFromPassword(date, cost)
	cvvHash, _ := bcrypt.GenerateFromPassword(cvv, cost)

	numberdb := string(numberHash)
	datedb := string(dateHash)
	cvvdb := string(cvvHash)

	fmt.Println("[Hashsum]")
	fmt.Println("----------------------------")
	fmt.Printf("[%s]\n", numberdb)
	fmt.Printf("[%s]\n", datedb)
	fmt.Printf("[%s]\n", cvvdb)
	fmt.Println("----------------------------")

	if err != nil {
		log.Printf("Failed unmarsheling: %s", err)
	}

	log.Print("your card added")
	return c.String(http.StatusOK, "successful request")
}
