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


func (s *SubCategoriesHandler) CreateSubCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categori_id := vars["categori_id"]
	CategoriID_int, err := strconv.Atoi(categori_id)
	if categori_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Categori ID not found in path",
		})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Invalid Categori ID format",
		})
		return
	}

	var create_res models.CreateCategoryRequest
	err = json.NewDecoder(r.Body).Decode(&create_res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.CategoriesRes{
			Status:      "error",
			Description: "Invalid request body",
		})
		return
	}

	create_subcategories, err := s.sercice.CreateSubCategory(&create_res.Name, &CategoriID_int)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.SubCategoriesRes{
			Status:      "error",
			Description: "Ошибка сервера",
		})
	}

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(models.SubCategoriesRes{
		Status:      create_subcategories,
		Description: "created",
	})
}