package app

import (
	handler "github.com/IDL13/echo/internal/handlers"
	"github.com/labstack/echo/v4"
)

type App struct {
	h    *handler.Handler
	r    *handler.RedisHandler
	a    *handler.Autorisation
	echo *echo.Echo
}
