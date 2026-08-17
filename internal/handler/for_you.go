package handler

import (
	"encoding/json"
	"net/http"

	"github.com/juraevibrahim01/jura/internal/service"
)

type TestHandler struct {
	service service.TestService
}

func NewForYouHandler(service service.TestService) *TestHandler {
	return &TestHandler{
		service: service,
	}
}

func (h *TestHandler) GetTest(w http.ResponseWriter, r *http.Request) {
	response := h.service.GetTest()

	result := map[string]interface{}{
		"data": response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)
}

func (h *TestHandler) GetTestNull(w http.ResponseWriter, r *http.Request) {
	response := h.service.GetTestNull()

	result := map[string]interface{}{
		"data": response,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)

}


