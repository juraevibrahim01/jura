package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/service"
)

type SubCategoriesHandler struct {
	sercice *service.SubCategoriesService
}

func NewSubCategories(sercice *service.SubCategoriesService) *SubCategoriesHandler {
	return &SubCategoriesHandler{sercice: sercice}
}

func (s *SubCategoriesHandler) GetSubCategories(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	categori_id := vars["categori_id"]
	if categori_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.SubCategoriesRes{
			Status:      "error",
			Description: "Categori ID not found in path",
		})
	}

	categori_id_int, err := strconv.Atoi(categori_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.SubCategoriesRes{
			Status:      "error",
			Description: "Invalid Categori ID format",
		})
	}

	subcategories, err := s.sercice.GetSubCategories(&categori_id_int)
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(models.SubCategoriesRes{
			Status:      "error",
			Description: "Ошибка сервера",
		})
	}

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(models.SubCategoriesRes{
		Status:        "success",
		Description:   "Подкатегории получены",
		SubCategories: *subcategories,
	})
}
