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

func NewDate() *Date {
	return &Date{}
}

func NewRedis() *Redis {
	return &Redis{}
}

type Redis struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type Name struct {
	Name string `json:"name"`
}

type Date struct {
	Email string `json:"email"`
	Msg   string `json:"msg"`
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

func (d *Date) Unmarshal(c echo.Context) *Date {

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
	}

	err = json.Unmarshal(b, &d)
	return d
}

func (r *Redis) Unmarshal(c echo.Context) *Redis {

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
	}

	err = json.Unmarshal(b, &r)
	return r
}
