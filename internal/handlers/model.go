package handler

import (
	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
)

type Handler struct {
	d *encryption.Date
	n *unmarshal.Name
}

type RedisHandler struct {
	r *unmarshal.Redis
}

type Autorisation struct {
	a *unmarshal.Auth
}
