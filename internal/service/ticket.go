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

func (s *Ticket_service) GetTickets(email *string) ([]models.Ticket, error) {
	return s.repository.GetTickets(email)
}
