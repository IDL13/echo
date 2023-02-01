package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IDL13/echo/internal/db"
	"github.com/IDL13/echo/internal/encryption"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	d        *encryption.Date
	database *db.Repository
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
	h.database = db.New()
	text := h.database.FindAll()
	log.Println(text)
	err := c.String(http.StatusOK, "[SUCCESSFUL ADDING]")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (h *Handler) AddOneHandler(c echo.Context) error {

	h.d = encryption.New()
	h.database = db.New()

	date, err := h.d.Encryption(c)
	if err != nil {
		fmt.Println("Error in Encryption")
	}
	err = h.database.Insert(date)
	if err != nil {
		log.Fatal("E-R-R-O-R")
	}

	return c.String(http.StatusOK, "successful request")
}

func (h *Handler) FindOneHandler(c echo.Context) error {

	h.d = encryption.New()
	h.database = db.New()

	err := h.database.FindOne("$2a$10$IpboAY23NFimUtPpeiY7h.fx4vXLMYSSVq9W.CNbqzUE6V30JPBPu")
	if err != nil {
		fmt.Println(err)
	}
	return c.String(http.StatusOK, "successful request")
}
