package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/juraevibrahim01/jura/internal/middleware"
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/service"
)

type AIHandler struct {
	service *service.AI_service
}

func NewAIHandler(service *service.AI_service) *AIHandler {
	return &AIHandler{service: service}
}

func (h *AIHandler) Chat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Проверяем авторизацию (из middleware)
	claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(models.ResponseAI{
			Status:      "error",
			Description: "Unauthorized",
		})
		return
	}

	// 2. Декодируем тело запроса
	var requestData models.RequestAI
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.ResponseAI{
			Status:      "error",
			Description: "Invalid request body",
		})
		return
	}

	// 3. Проверяем, что вопрос не пустой
	if requestData.Question == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.ResponseAI{
			Status:      "error",
			Description: "Question cannot be empty",
		})
		return
	}

	// 4. Отправляем запрос в AI сервис
	reply, err := h.service.SendMessage(&claims.UserID, requestData.Question)
	if err != nil {
		// Обрабатываем разные типы ошибок по содержимому
		statusCode := http.StatusInternalServerError
		description := "Ошибка при обработке запроса"

		if strings.Contains(err.Error(), "no response") {
			statusCode = http.StatusServiceUnavailable
			description = "Сервис AI временно недоступен"
		} else if strings.Contains(err.Error(), "Invalid API key") {
			statusCode = http.StatusInternalServerError
			description = "Ошибка конфигурации сервиса"
		}

		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(models.ResponseAI{
			Status:      "error",
			Description: description,
		})
		return
	}

	// 5. Успешный ответ
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.ResponseAI{
		Status:      "success",
		Description: "Ответ получен",
		Answer:      reply,
	})
}
