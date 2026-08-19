package service

import (
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/repository"
)

type Ticket_service struct {
	repository *repository.Ticket_repository
}

func Ticket_new_service(repository *repository.Ticket_repository) *Ticket_service {
	return &Ticket_service{repository: repository}
}

func (s *Ticket_service) GetTickets(userID, projectID *int) ([]models.Ticket, error) {
	return s.repository.GetTickets(userID, projectID)
}

func (s *Ticket_service) Ticket_create(data, name, module, precondition, steps, expectation_res, actual_res, comment *string, user_id, project_id *int) error {
	return s.repository.Ticket_create(data, name, module, precondition, steps, expectation_res, actual_res, comment, user_id, project_id)
}

func (s *Ticket_service) GetTicketsByID(userID, projectID, ticketsID *int) ([]models.Ticket, error) {
	return s.repository.GetTicketsByID(userID, projectID, ticketsID)
}
