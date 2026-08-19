package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	// claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	// if !ok || claims == nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	_ = json.NewEncoder(w).Encode(models.TicketsResponse{
	// 		Status:      "error",
	// 		Description: "Unauthorized",
	// 	})
	// 	return
	// }

	vars := mux.Vars(r)

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

	projectID := vars["project_id"]
	if projectID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Project ID not found in path",
		})
		return
	}

	projectID_int, err := strconv.Atoi(projectID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TestKeysResponse{
			Status:      "error",
			Description: "Invalid project_id format",
		})
		return
	}

	tickets, err := h.service.GetTickets(&UserIDInt, &projectID_int)
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

func (h *Ticket_handler) GetTicketsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	// if !ok || claims == nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	_ = json.NewEncoder(w).Encode(models.TicketsResponse{
	// 		Status:      "error",
	// 		Description: "Unauthorized",
	// 	})
	// 	return
	// }

	vars := mux.Vars(r)

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

	projectID := vars["project_id"]
	if projectID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Project ID not found in path",
		})
		return
	}
	projectID_int, err := strconv.Atoi(projectID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Invalid project_id format",
		})
		return
	}

	ticketsID := vars["id"]
	if ticketsID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Tickets ID not found in path",
		})
	}
	ticketsID_int, err := strconv.Atoi(ticketsID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Invalid tickets_id format",
		})
		return
	}

	tickets, err := h.service.GetTicketsByID(&UserIDInt, &projectID_int, &ticketsID_int)
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

func (h *Ticket_handler) Ticket_create(w http.ResponseWriter, r *http.Request) {

	// claims, ok := r.Context().Value(middleware.ClaimsKey).(*models.Claims)
	// if !ok || claims == nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	_ = json.NewEncoder(w).Encode(models.TicketsResponse{
	// 		Status:      "error",
	// 		Description: "Unauthorized",
	// 	})
	// 	return
	// }

	vars := mux.Vars(r)

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

	projectID := vars["project_id"]
	if projectID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Project ID not found in path",
		})
		return
	}
	projectID_int, err := strconv.Atoi(projectID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Invalid project_id format",
		})
		return
	}

	var request models.TicketCreateRequest

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Ошибка при разборе запроса",
		})
		return
	}

	err = h.service.Ticket_create(&request.Data, &request.Name, &request.Module, &request.Precondition, &request.Steps, &request.ExpectationRes, &request.ActualRes, &request.Comment, &UserIDInt, &projectID_int)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.TicketsResponse{
			Status:      "error",
			Description: "Ошибка сервера при создании тикета",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(models.TicketsResponse{
		Status:      "success",
		Description: "Тикет успешно создан",
	})
}
