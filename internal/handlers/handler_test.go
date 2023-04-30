package handler

import (
	"bytes"
	"context"
	"net/http/httptest"
	"testing"

	"github.com/IDL13/echo/internal/encryption"
	mock_serv "github.com/IDL13/echo/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestInsertOneHandler(t *testing.T) {
	type mockBehaviour func(m *mock_serv.MockRepository, card *encryption.Date)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		date          encryption.Date
		input         string
		status        int
	}{
		{
			name: "OK",
			mockBehaviour: func(m *mock_serv.MockRepository, card *encryption.Date) {
				m.EXPECT().Insert(context.TODO(), card).Return(nil)
			},
			date: encryption.Date{
				Number: "Adolf",
				Date:   "12.12",
				CVV:    "L8F?N5goKvJ90MZl$2a$10$/E7bXIWqtxQoHVxByLWjvO7VJxPEOcsQreMeY.a9I35/Ih6QVxxzy8SPea4dzvL_g0XFK",
			},
			input:  `{Name:"Adolf", Date:"12.12", CVV:"453"}`,
			status: 400,
		},

		{
			name: "NO",
			mockBehaviour: func(m *mock_serv.MockRepository, card *encryption.Date) {
				m.EXPECT().Insert(context.TODO(), card).Return(nil)
			},
			date: encryption.Date{
				Number: "Adolf",
				Date:   "12.12",
				CVV:    "L8F?N5goKvJ90MZl$2a$10$/E7bXIWqtxQoHVxByLWjvO7VJxPEOcsQreMeY.a9I35/Ih6QVxxzy8SPea4dzvL_g0XFK",
			},
			input:  `{Name:"Adolf", Date:"12.12", CVV:"453"}`,
			status: 200,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			controler := gomock.NewController(t)
			defer controler.Finish()

			auth := mock_serv.NewMockRepository(controler)

			test.mockBehaviour(auth, &test.date)

			//Создание тестового сервера
			e := echo.New()
			h := New()
			e.POST("/test", h.AddOneHandler)

			// //Отправка тестового POST запроса
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(test.input))

			e.ServeHTTP(w, req)

			//Проверка на соответствие
			assert.Equal(t, w.Code, test.status)
		})
	}

}
