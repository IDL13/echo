package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/IDL13/echo/internal/db"
	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
	"github.com/IDL13/echo/pkg/api"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
	d *encryption.Date
	n *unmarshal.Name
}

func (h *Handler) StartHandler(c echo.Context) error {
	err := c.String(http.StatusOK, "[SERVER STARTED]")
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) SmtpHandler(c echo.Context) error {

	date := unmarshal.NewDate()
	date.Unmarshal(c)

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	genConn := api.NewAdderClient(conn)
	res, err := genConn.Add(context.TODO(), &api.AddRequest{Email: date.Email, Msg: date.Msg})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
	return nil
}

func (h *Handler) FindAllHandler(c echo.Context) error {

	r := db.New()

	text, err := r.FindAll(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range text {
		if err := c.String(http.StatusOK, val.Number); err != nil {
			log.Println(err)
		}
		if err := c.String(http.StatusOK, "\n"); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (h *Handler) AddOneHandler(c echo.Context) error {
	r := db.New()

	h.d = encryption.New()
	h.n = unmarshal.New()

	date, err := h.d.Encryption(c)

	if err != nil {
		log.Fatal(err)
	}

	go r.Insert(context.TODO(), date)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return c.String(http.StatusOK, "successful request")
}

func (h *Handler) FindOneHandler(c echo.Context) error {
	r := db.New()

	h.n = unmarshal.New()

	name := h.n.Unmarshal(c)

	go r.FindOne(context.TODO(), name)
	// if err != nil {
	// 	log.Println(err)
	// }
	return c.String(http.StatusOK, "successful request")
}
