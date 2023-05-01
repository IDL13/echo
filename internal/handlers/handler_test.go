package handler

import (
	"context"
	"testing"

	"github.com/IDL13/echo/internal/encryption"
	mock_serv "github.com/IDL13/echo/internal/mock"
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

// func TestFindAll(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	repo := mock_serv.NewMockRepository(ctrl)

// 	ctx := context.TODO()

// 	mas := []

// 	repo.EXPECT().FindAll(ctx).Return(, nil )

// 	text, err := repo.FindAll(ctx)
// }

// func TestInsertOneHandler(t *testing.T) {
// 	type mockBehaviour func(m *mock_serv.MockRepository, card *encryption.Date)

// 	tests := []struct {
// 		name          string
// 		mockBehaviour mockBehaviour
// 		date          *encryption.Date
// 		input         string
// 		status        int
// 	}{
// 		{
// 			name: "OK",
// 			mockBehaviour: func(m *mock_serv.MockRepository, card *encryption.Date) {
// 				m.EXPECT().Insert(context.TODO(), card).Times(1).Return(nil)
// 			},
// 			date: &encryption.Date{
// 				Number: "Adolf",
// 				Date:   "12.12",
// 				CVV:    "453",
// 			},
// 			input: `{
// 				"number":"Nabokov", "date":"12.12", "CVV":"453"
// 			}`,
// 			status: 400,
// 		},

// 		{
// 			name: "NO",
// 			mockBehaviour: func(m *mock_serv.MockRepository, card *encryption.Date) {
// 				m.EXPECT().Insert(context.TODO(), card).Return(nil)
// 			},
// 			date: &encryption.Date{
// 				Number: "Adolf",
// 				Date:   "12.12",
// 				CVV:    "L8F?N5goKvJ90MZl$2a$10$/E7bXIWqtxQoHVxByLWjvO7VJxPEOcsQreMeY.a9I35/Ih6QVxxzy8SPea4dzvL_g0XFK",
// 			},
// 			input: `{
// 				"number":"Nabokov", "date":"12.12", "CVV":"453"
// 			}`,
// 			status: 200,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			controler := gomock.NewController(t)
// 			defer controler.Finish()

// 			auth := mock_serv.NewMockRepository(controler)

// 			test.mockBehaviour(auth, test.date)

// 			if test.status != 400 {
// 				t.Error("not 400")
// 			}

// 			// //Создание тестового сервера
// 			// e := echo.New()
// 			// h := New()
// 			// e.POST("/test", h.AddOneHandler)

// 			// // //Отправка тестового POST запроса
// 			// w := httptest.NewRecorder()
// 			// req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(test.input))

// 			// e.ServeHTTP(w, req)

// 			//Проверка на соответствие
// 			// assert.Equal(t, 400, test.status)
// 		})
// 	}

// }
