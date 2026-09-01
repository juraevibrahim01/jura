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

func (r *Ticket_repository) GetTickets(projectID, categoryID, subcategoryID *int) ([]models.Ticket, error) {
	query := `
		SELECT
			t.id,
			t."title"
		FROM tickets t
		JOIN subcategories sc on sc.id = t.subcategory_id
		JOIN categories c on c.id = sc.categori_id
		JOIN projects p on p.id = c.project_id
		where p.id = $1 and c.id = $2 and sc.id = $3;
	`

	rows, err := r.postgres.DB.Query(query, projectID, categoryID, subcategoryID)
	if err != nil {
		log.Print("Ошибка при получении тикетов: ", err)
		return nil, err
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		err = rows.Scan(&ticket.ID, &ticket.Title)
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

func (r *Ticket_repository) GetTicketsByID(userID, projectID, ticketsID *int) (*models.Ticket, error) {
	query := `
		SELECT t.id, t."title", t."priority", t."severity", t."environment", t."steps", t."expected_result", t."actual_result", t."attachments", t."created_at"
		FROM tickets t
        where t.user_id = $1 and t.project_id = $2 and t.id = $3;
	`

	row := r.postgres.DB.QueryRow(query, userID, projectID, ticketsID)

	var ticket models.Ticket

	if err := row.Scan(&ticket.ID, &ticket.Title, &ticket.Priority, &ticket.Severity, &ticket.Environment, &ticket.Steps, &ticket.ExpectedResult, &ticket.ActualResult, &ticket.Attachments, &ticket.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Print("Ошибка при получении багрепорта: ", err)
		return nil, err
	}

	return &ticket, nil
}

func (r *Ticket_repository) Ticket_create(title, priority, severity, environment, steps, expected_res, actual_res, attachments *string, subCategoryID_int *int) error {

	query := `
		INSERT INTO tickets ("title", "priority", "severity", "environment", "steps", "expected_result", "actual_result", "attachments", subcategory_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := r.postgres.DB.Exec(query, title, priority, severity, environment, steps, expected_res, actual_res, attachments, subCategoryID_int)
	if err != nil {
		log.Print("Ошибка при создании тикета: ", err)
		return err
	}
	return nil
}
