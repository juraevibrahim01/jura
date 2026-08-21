package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	// claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	// if !ok || claims == nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
	// 		Status:      "error",
	// 		Description: "Unauthorized",
	// 	})
	// 	return
	// }

	vars := mux.Vars(r)
	projectID := vars["project_id"]
	ProjectID_int, err := strconv.Atoi(projectID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Invalid Project ID format",
		})
		return
	}

	if projectID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Project ID not found in path",
		})
		return
	}

	var UserID string
	UserID = r.Header.Get("X-User-UserID")
	if UserID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "UserID not found in header",
		})
		return
	}

	UserIDInt, err := strconv.Atoi(UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Invalid UserID format",
		})
		return
	}

	testKeys, err := h.service.GetTestKeys(&UserIDInt, &ProjectID_int)
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

	// claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	// if !ok || claims == nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
	// 		Status:      "error",
	// 		Description: "Unauthorized",
	// 	})
	// 	return
	// }

	var UserID string
	UserID = r.Header.Get("X-User-UserID")
	if UserID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "UserID not found in header",
		})
		return
	}
	UserIDInt, err := strconv.Atoi(UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Invalid UserID format",
		})
		return
	}

	vars := mux.Vars(r)
	_ = vars["project_id"]
	idStr := vars["id"]
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Не указан id тестового кейса",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Неверный id тестового кейса",
		})
		return
	}

	testKey, err := h.service.GetTestKeyByID(&id, &UserIDInt)
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

	// claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	// if !ok || claims == nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
	// 		Status:      "error",
	// 		Description: "Unauthorized",
	// 	})
	// 	return
	// }

	var UserID string
	UserID = r.Header.Get("X-User-UserID")
	userIDInt, err := strconv.Atoi(UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Неверный id пользователя",
		})
		return
	}

	vars := mux.Vars(r)

	projectID := vars["project_id"]

	projectIDInt, err := strconv.Atoi(projectID)
	if projectID == "" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeyResponse{
			Status:      "error",
			Description: "Неверный id проекта",
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

	if err := h.service.CreateTestKey(&request, &userIDInt, &projectIDInt); err != nil {
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

func (h *Test_keys_handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projects, err := h.service.GetProjects()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(struct {
			Status      string           `json:"status"`
			Description string           `json:"description"`
			Projects    []models.Project `json:"projects"`
		}{
			Status:      "error",
			Description: "Ошибка сервера",
			Projects:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct {
		Status      string           `json:"status"`
		Description string           `json:"description"`
		Projects    []models.Project `json:"projects"`
	}{
		Status:      "success",
		Description: "Проекты получены",
		Projects:    projects,
	})
}

func (h *Test_keys_handler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(struct {
			Status      string          `json:"status"`
			Description string          `json:"description"`
			Project     *models.Project `json:"project,omitempty"`
		}{
			Status:      "error",
			Description: "Не указан id проекта",
			Project:     nil,
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(struct {
			Status      string            `json:"status"`
			Description string            `json:"description"`
			Project     *models.ProjectID `json:"project,omitempty"`
		}{
			Status:      "error",
			Description: "Неверный id проекта",
			Project:     nil,
		})
		return
	}

	project, err := h.service.GetProjectByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(struct {
			Status      string            `json:"status"`
			Description string            `json:"description"`
			Project     *models.ProjectID `json:"project,omitempty"`
		}{
			Status:      "error",
			Description: "Ошибка сервера",
			Project:     nil,
		})
		return
	}

	if project == nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(struct {
			Status      string            `json:"status"`
			Description string            `json:"description"`
			Project     *models.ProjectID `json:"project,omitempty"`
		}{
			Status:      "error",
			Description: "Проект не найден",
			Project:     nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct {
		Status      string            `json:"status"`
		Description string            `json:"description"`
		Project     *models.ProjectID `json:"project,omitempty"`
	}{
		Status:      "success",
		Description: "Проект получен",
		Project:     project,
	})
}
