package handler

import (
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

func (h *Handler) GetHandler(c echo.Context) error {
	err := c.String(http.StatusOK, "[SERVER STARTED]")
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) PostHandler(c echo.Context) error {

	h.d = encryption.New()
	h.database = db.New()

	h.d.Encryption(c)
	err := h.database.Insert()
	if err != nil {
		log.Fatal("E-R-R-O-R")
	}

	return c.String(http.StatusOK, "successful request")
}
