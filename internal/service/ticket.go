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

func (s *Ticket_service) GetTickets(projectID, categoryID, subcategoryID *int) ([]models.Ticket, error) {
	return s.repository.GetTickets(projectID, categoryID, subcategoryID)
}

func (s *Ticket_service) Ticket_create(title, priority, severity, environment, steps, expected_res, actual_res, attachments *string, subCategoryID_int *int) error {
	return s.repository.Ticket_create(title, priority, severity, environment, steps, expected_res, actual_res, attachments, subCategoryID_int)
}

func (s *Ticket_service) GetTicketsByID(userID, projectID, ticketsID *int) (*models.Ticket, error) {
	return s.repository.GetTicketsByID(userID, projectID, ticketsID)
}
