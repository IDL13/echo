package encryption

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func New() *Date {
	return &Date{}
}

type Date struct {
	Number string `json:"number"`
	Date   string `json:"date"`
	CVV    string `json:"CVV"`
}

func (d *Date) Encryption(c echo.Context) (*Date, error) {
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
	}

	err = json.Unmarshal(b, &d)

	// Convert Number, Date, CVV
	number := []byte(d.Number)
	date := []byte(d.Date)
	cvv := []byte(d.CVV)

	cost := 10

	numberHash, _ := bcrypt.GenerateFromPassword(number, cost)
	dateHash, _ := bcrypt.GenerateFromPassword(date, cost)
	cvvHash, _ := bcrypt.GenerateFromPassword(cvv, cost)

	d.Number = string(numberHash)
	d.Date = string(dateHash)
	d.CVV = string(cvvHash)

	if err != nil {
		log.Printf("Failed unmarsheling: %s", err)
	}

	return d, nil
}
