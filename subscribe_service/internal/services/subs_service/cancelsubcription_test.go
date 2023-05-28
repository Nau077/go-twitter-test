package subs_service

import (
	"encoding/json"
	"errors"
	mock_repository "go_subs_service/internal/services/subs_service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSubsService_cancelSub(t *testing.T) {
	testCases := []struct {
		name           string
		expectedParam  string
		expectedMsg    string
		expectedErr    error
		expectedStatus int
	}{
		{
			name:           "Successful cancellation",
			expectedParam:  "test_user",
			expectedMsg:    "Subscription canceled",
			expectedErr:    nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Cancellation error",
			expectedParam:  "test_user",
			expectedMsg:    "",
			expectedErr:    errors.New("some error"),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Создаем мок-объект для SubcribeRepository
			mockRepo := mock_repository.NewMockSubcribeRepository(ctrl)

			// Создаем сервис с мок-репозиторием
			service := &Service{subscriptionRepository: mockRepo}

			if tc.expectedErr == nil {
				// Устанавливаем ожидание вызова метода CancelSubscription и настраиваем его результаты
				mockRepo.EXPECT().CancelSubcription(gomock.Any(), tc.expectedParam).Return(tc.expectedMsg, nil)
			} else {
				mockRepo.EXPECT().CancelSubcription(gomock.Any(), tc.expectedParam).Return("", tc.expectedErr)
			}

			// Создаем тестовый запрос
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Устанавливаем параметр "user" в запросе
			c.Request, _ = http.NewRequest(http.MethodGet, "/path/to/endpoint?user="+tc.expectedParam, nil)

			// Вызываем функцию-обработчик CancelSub
			handler := service.CancelSub(nil)
			handler(c)

			// Проверяем ответ и статус код
			assert.Equal(t, tc.expectedStatus, w.Code)

			if tc.expectedErr == nil {
				var response struct {
					Result string `json:"result"`
				}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Проверяем ожидаемое сообщение
				assert.Equal(t, tc.expectedMsg, response.Result)
			} else {
				var response struct {
					Error string `json:"error"`
				}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				// Проверяем ожидаемое сообщение об ошибке
				assert.Equal(t, tc.expectedErr.Error(), response.Error)
			}
		})
	}
}
