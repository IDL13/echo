package handler

import (
	"log"
	"net/http"

	"github.com/IDL13/echo/internal/db"
	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	d        *encryption.Date
	database *db.Repository
	n        *unmarshal.Name
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) StartHandler(c echo.Context) error {
	err := c.String(http.StatusOK, "[SERVER STARTED]")
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) FindAllHandler(c echo.Context) error {
	text := h.database.FindAll()
	for _, val := range text {
		if err := c.String(http.StatusOK, val); err != nil {
			log.Println(err)
		}
		if err := c.String(http.StatusOK, "\n"); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (h *Handler) AddOneHandler(c echo.Context) error {

	h.d = encryption.New()
	h.database = db.New()
	h.n = unmarshal.New()

	date, err := h.d.Encryption(c)
	if err != nil {
		log.Println("Error in Encryption")
	}

	err = h.database.Insert(date)
	if err != nil {
		log.Fatal("Error from database INSERT")
	}

	return c.String(http.StatusOK, "successful request")
}

func (h *Handler) FindOneHandler(c echo.Context) error {

	h.n = unmarshal.New()

	name := h.n.Unmarshal(c)

	err := h.database.FindOne(name)
	if err != nil {
		log.Println(err)
	}
	return c.String(http.StatusOK, "successful request")
}
