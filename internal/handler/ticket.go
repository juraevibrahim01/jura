package handler

import (
	"encoding/json"
	"net/http"

	"github.com/juraevibrahim01/jura/internal/middleware"
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/service"
)

type Ticket_handler struct {
	service *service.Ticket_service
}

func Ticket_new_handler(service *service.Ticket_service) *Ticket_handler {
	return &Ticket_handler{service: service}
}

func (h *Ticket_handler) GetTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	if !ok || claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Unauthorized",
		})
		return
	}

	tickets, err := h.service.GetTickets(&claims.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Ошибка сервера",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.TicketsResponse{
		Status:      "success",
		Description: "Тикеты получены",
		Tickets:     tickets,
	})
}
