package app

import (
	"fmt"

	handler "github.com/IDL13/echo/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	h    *handler.Handler
	r    *handler.RedisHandler
	a    *handler.Autorisation
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.h = handler.New()
	a.r = handler.NewRedis()
	a.echo = echo.New()

	a.echo.GET("/", a.h.StartHandler)
	a.echo.GET("/findAll", a.h.FindAllHandler)

	a.echo.GET("/redis/:key", a.r.GetHandler)

	a.echo.POST("/addOne", a.h.AddOneHandler)
	a.echo.POST("/findOne", a.h.FindOneHandler)
	a.echo.POST("/smtp", a.h.SmtpHandler)
	a.echo.POST("/autorisation", a.a.AuthHandler)
	a.echo.POST("/registration", a.a.RegHandler)

	a.echo.POST("/redis", a.r.SetHandler)

	a.echo.DELETE("delete/:id", a.h.DeleteHandler)

	a.echo.PUT("put/:id", a.h.PutHandler)

	DefaulCorsConfig := middleware.DefaultCORSConfig

	a.echo.Use(middleware.CORSWithConfig(DefaulCorsConfig))

	return a, nil
}

func (a *App) Run() {
	fmt.Println("[SERVER STARTED]")
	a.echo.Logger.Fatal(a.echo.Start(":8080"))
}
