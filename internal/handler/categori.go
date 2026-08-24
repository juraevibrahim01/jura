package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/service"
)

type CategoryHandler struct {
	service *service.CategotiService
}

func NewCategory(service *service.CategotiService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (c *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	project_id := vars["project_id"]
	ProjectID_int, err := strconv.Atoi(project_id)
	if project_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Project ID not found in path",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Invalid Project ID format",
		})
		return
	}

	categories, err := c.service.GetCategories(&ProjectID_int)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.CategoriesRes{
			Status:      "error",
			Description: "Ошибка сервера",
		})
	}

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(models.CategoriesRes{
		Status:     "success",
		Categories: *categories,
	})
}
