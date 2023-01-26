package handler

import (
	"net/http"

	"github.com/IDL13/echo/internal/encryption"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	d *encryption.Date
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

	h.d.Encryption(c)

	return c.String(http.StatusOK, "successful request")
}
