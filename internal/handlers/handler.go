package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/IDL13/echo/internal/db"
	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
	"github.com/IDL13/echo/pkg/api"
	"github.com/IDL13/echo/pkg/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

const (
	tokenTTL = 12 * time.Hour
)

func New() *Handler {
	return &Handler{}
}

func NewRedis() *RedisHandler {
	return &RedisHandler{}
}

func NewAuthorisation() *Autorisation {
	return &Autorisation{}
}

func (a *Autorisation) AuthHandler(c echo.Context) error {

	user := a.a.UnmarshalFoAuth(c)

	r := db.New()

	flag, id := r.FindOneById(context.TODO(), user)

	if flag == 1 {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		})

		// err := godotenv.Load()
		// if err != nil {
		// 	log.Println(".env is empty")
		// }

		encryption.LoadDotenv()

		token.SignedString([]byte(os.Getenv("SALT")))

		conn := redis.Connection()

		ctx := context.TODO()

		err := conn.Set(ctx, strconv.Itoa(id), token, 0).Err()
		if err != nil {
			log.Println("Error from Redi set")
		}
		return c.String(http.StatusOK, "successful request")
	} else {
		return c.String(http.StatusOK, "Bad request")
	}
}

func (a *Autorisation) RegHandler(c echo.Context) error {

	user := a.a.UnmarshalFoReg(c)

	r := db.New()

	err := r.InsertOneUser(context.TODO(), user)

	if err != nil {
		fmt.Println("user not registration")
	}

	return c.String(http.StatusOK, "successful request")
}

func (r *RedisHandler) SetHandler(c echo.Context) error {
	conn := redis.Connection()

	ctx := context.TODO()

	r.r.Unmarshal(c)

	err := conn.Set(ctx, r.r.Key, r.r.Val, 0).Err()
	if err != nil {
		log.Println("Error from Redi set")
	}

	return nil
}

func (r *RedisHandler) GetHandler(c echo.Context) error {
	key := c.Param("key")

	conn := redis.Connection()

	ctx := context.TODO()

	val, err := conn.Get(ctx, key).Result()
	if err != nil {
		log.Fatal("Error from redis GetHandler")
	}

	fmt.Println(val)

	return nil
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

func (h *Handler) DeleteHandler(c echo.Context) error {
	r := db.New()

	id, _ := strconv.Atoi(c.Param("id"))

	err := r.Delete(context.TODO(), id)
	if err != nil {
		log.Println(err)
	}
	return c.String(http.StatusOK, "successful request")
}

func (h *Handler) PutHandler(c echo.Context) error {
	r := db.New()

	h.d = encryption.New()
	h.n = unmarshal.New()

	date, err := h.d.Encryption(c)

	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err = r.Put(context.TODO(), id, date)
	if err != nil {
		log.Println(err)
	}
	return c.String(http.StatusOK, "successful request")
}
