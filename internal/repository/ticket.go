package repository

import (
	"database/sql"
	"log"

	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/pkg"
)

type Ticket_repository struct {
	postgres *pkg.Postgres
}

func Ticket_new_repository(postgres *pkg.Postgres) *Ticket_repository {
	return &Ticket_repository{postgres: postgres}
}

func (r *Ticket_repository) GetTickets(userID, projectID *int) ([]models.Ticket, error) {
	query := `
		SELECT t.id, t.title
		FROM tickets t
        where t.user_id = $1 and t.project_id = $2 and t.id = $3;
	`

	rows, err := r.postgres.DB.Query(query, userID, projectID)
	if err != nil {
		log.Print("Ошибка при получении тикетов: ", err)
		return nil, err
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		err = rows.Scan(&ticket.ID, &ticket.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return []models.Ticket{}, nil
			}
			log.Print("Ошибка при сканировании тикетов: ", err)
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *Ticket_repository) GetTicketsByID(userID, projectID, ticketsID *int) ([]models.Ticket, error) {
	query := `
		SELECT t.id, t.data, t.title, t.module, t.precondition, t.steps, t.expectation_res, t.actual_res, t.comment
		FROM tickets t
        where t.user_id = $1 and t.project_id = $2 and t.id = $3;
	`

	rows, err := r.postgres.DB.Query(query, userID, projectID, ticketsID)
	if err != nil {
		log.Print("Ошибка при получении тикетов: ", err)
		return nil, err
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		err = rows.Scan(&ticket.ID, &ticket.Data, &ticket.Name, &ticket.Module, &ticket.Precondition, &ticket.Steps, &ticket.Expectation_res, &ticket.Actual_res, &ticket.Comment)
		if err != nil {
			if err == sql.ErrNoRows {
				return []models.Ticket{}, nil
			}
			log.Print("Ошибка при сканировании тикетов: ", err)
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *Ticket_repository) Ticket_create(data, name, module, precondition, steps, expectation_res, actual_res, comment *string, user_id, project_id *int) error {

	query := `
		INSERT INTO tickets (date, "title", "module", "precondition", "steps", "expectation_res", "actual_res", "comment", user_id, project_id)
		VALUES ($1, (SELECT id FROM users WHERE email = $2), $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := r.postgres.DB.Exec(query, data, name, module, precondition, steps, expectation_res, actual_res, comment, user_id, project_id)
	if err != nil {
		log.Print("Ошибка при создании тикета: ", err)
		return err
	}
	return nil
}
