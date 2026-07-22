package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/juraevibrahim01/jura/internal/middleware"
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/service"
)

type Test_keys_handler struct {
	service *service.Test_keys_service
}

func New_Test_keys_handler(service *service.Test_keys_service) *Test_keys_handler {
	return &Test_keys_handler{service: service}
}

func (h *Test_keys_handler) GetTestKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Unauthorized",
		})
		return
	}

	testKeys, err := h.service.GetTestKeys(&claims.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Ошибка сервера",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
		Status:      "success",
		Description: "Тестовые ключи получены",
		TestKeys:    testKeys,
	})
}

func (h *Test_keys_handler) GetTestKeyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Unauthorized",
		})
		return
	}

	idPart := strings.TrimPrefix(r.URL.Path, "/test-keys/")
	if idPart == "" || idPart == "/" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Не указан id тестового кейса",
		})
		return
	}

	id, err := strconv.Atoi(idPart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Неверный id тестового кейса",
		})
		return
	}

	testKey, err := h.service.GetTestKeyByID(&id, &claims.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Ошибка сервера",
		})
		return
	}

	if testKey == nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Тестовый кейс не найден",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
		Status:      "success",
		Description: "Тестовый кейс получен",
		TestKey:     testKey,
	})
}

func (h *Test_keys_handler) CreateTestKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Unauthorized",
		})
		return
	}

	var request models.TestKeyCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Ошибка при разборе запроса",
		})
		return
	}

	if err := h.service.CreateTestKey(&request, &claims.UserID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Ошибка сервера при создании тестового кейса",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
		Status:      "success",
		Description: "Тестовый кейс успешно создан",
	})
}
