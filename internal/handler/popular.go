package handler

import (
	"encoding/json"
	"net/http"

	"github.com/juraevibrahim01/jura/internal/service"
)

type PopularHandler struct {
	service service.PopularService
}

func NewPopularHandler(service service.PopularService) *PopularHandler {
	return &PopularHandler{
		service: service,
	}
}

func (h *PopularHandler) PopularGetTest(w http.ResponseWriter, r *http.Request) {
	response := h.service.GetPopular()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (h *PopularHandler) PopularGetTestNull(w http.ResponseWriter, r *http.Request) {
	response := h.service.GetPopularNull()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
