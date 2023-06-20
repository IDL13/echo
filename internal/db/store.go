package db

import (
	"context"

	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
)

// generate mock repository
//
//go:generate mockgen -source=store.go -destination=../mock/mock.go -package=mock_serv
type Repository interface {
	Insert(ctx context.Context, card *encryption.Date) error
	FindAll(ctx context.Context) (mas []Card, err error)
	FindOne(ctx context.Context, number *unmarshal.Name) error
	Delete(ctx context.Context, id int) error
	Put(ctc context.Context, id int) error
}
