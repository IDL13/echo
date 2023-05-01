package handler

import (
	"context"
	"testing"

	"github.com/IDL13/echo/internal/db"
	"github.com/IDL13/echo/internal/encryption"
	mock_serv "github.com/IDL13/echo/internal/mock"
	"github.com/IDL13/echo/internal/unmarshal"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddOneHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_serv.NewMockRepository(ctrl)

	ctx := context.TODO()

	encription := &encryption.Date{
		Number: "Adolf",
		Date:   "12.12",
		CVV:    "L8F?N5goKvJ90MZl$2a$10$/E7bXIWqtxQoHVxByLWjvO7VJxPEOcsQreMeY.a9I35/Ih6QVxxzy8SPea4dzvL_g0XFK",
	}

	repo.EXPECT().Insert(ctx, encription).Return(nil)

	err := repo.Insert(ctx, encription)

	require.NoError(t, err)
}

func TestFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_serv.NewMockRepository(ctrl)

	ctx := context.TODO()

	var text []db.Card

	repo.EXPECT().FindAll(ctx).Return(text, nil)

	text, err := repo.FindAll(ctx)

	require.NoError(t, err)
}

func TestFindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_serv.NewMockRepository(ctrl)

	ctx := context.TODO()

	text := &unmarshal.Name{Name: "Nabokov"}

	repo.EXPECT().FindOne(ctx, text).Return(nil)

	err := repo.FindOne(ctx, text)

	require.NoError(t, err)
}
