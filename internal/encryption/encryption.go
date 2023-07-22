package encryption

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const SALT1 string = "L8F?N5goKvJ90MZl"
const SALT2 string = "8SPea4dzvL_g0XFK"

func New() *Date {
	return &Date{}
}

func (d *Date) Encryption(c echo.Context) (*Date, error) {
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
	}

	err = json.Unmarshal(b, &d)

	// Convert Number, Date, CVV
	cvv := []byte(d.CVV)

	cost := 10

	cvvHash, _ := bcrypt.GenerateFromPassword(cvv, cost)

	d.CVV = SALT1 + string(cvvHash) + SALT2

	if err != nil {
		log.Printf("Failed unmarsheling: %s", err)
	}

	return d, nil
}
