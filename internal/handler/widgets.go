package handler

import (
	"encoding/json"
	"net/http"

	"github.com/juraevibrahim01/jura/internal/service"
)

type WidgetsHandler struct {
	service service.WidgetsService
}

func NewWidgetsHandler(service service.WidgetsService) *WidgetsHandler {
	return &WidgetsHandler{service: service}
}

func (h *WidgetsHandler) GetWidgets(w http.ResponseWriter, r *http.Request) {
	response := h.service.GetWidgets()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
