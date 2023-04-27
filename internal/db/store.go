package db

import (
	"context"

	"github.com/IDL13/echo/internal/encryption"
	"github.com/IDL13/echo/internal/unmarshal"
)

type Repository interface {
	Insert(ctx context.Context, card *encryption.Date) error
	FindAll(ctx context.Context) (mas []Card, err error)
	FindOne(ctx context.Context, number *unmarshal.Name) error
}
