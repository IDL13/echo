package encryption

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func New() *Date {
	return &Date{}
}

func LoadDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is empty")
	}
}

func (d *Date) Encryption(c echo.Context) (*Date, error) {
	LoadDotenv()

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

	d.CVV = os.Getenv("SALT1") + string(cvvHash) + os.Getenv("SALT2")

	if err != nil {
		log.Printf("Failed unmarsheling: %s", err)
	}

	return d, nil
}

func PassHashing(pass string) string {
	// Convert Password
	pas := []byte(pass)

	cost := 10

	passwordHash, _ := bcrypt.GenerateFromPassword(pas, cost)

	NewPass := string(passwordHash)

	return NewPass
}

func CheckPassword(hashPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	return err
}
