package unmarshal

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
)

func New() *Name {
	return &Name{}
}

type Name struct {
	Name string `json:"name"`
}

func (n *Name) Unmarshal(c echo.Context) *Name {

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
	}

	err = json.Unmarshal(b, &n)
	if err != nil {
		log.Println(err)
	}

	return n
}
