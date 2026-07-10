package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/pkg"
)

type Ticket_repository struct {
	postgres *pkg.Postgres
}

func Ticket_new_repository(postgres *pkg.Postgres) *Ticket_repository {
	return &Ticket_repository{postgres: postgres}
}

func (r *Ticket_repository) GetTickets(email *string) ([]models.Ticket, error) {
	query := `
		SELECT t.id, t."Title"
		FROM tickets t
        JOIN users u on u.id = t.user_id
        WHERE u.email = $1;
	`

	rows, err := r.postgres.DB.Query(query, email)
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

func (r *Ticket_repository) Ticket_create(title *string, email *string, priority *string, severity *string, environment *string, stepsToReproduce *string, expectedResult *string, actualResult *string, attachments *[]string) error {

	attachmentsStr := strings.Join(*attachments, ",")

	query := `
		INSERT INTO tickets ("Title", user_id, "Priority", "Severity", "Environment", "Steps", "Expected_Result", "Actual_Result", "Attachments")
		VALUES ($1, (SELECT id FROM users WHERE email = $2), $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := r.postgres.DB.Exec(query, title, email, priority, severity, environment, stepsToReproduce, expectedResult, actualResult, attachmentsStr)
	if err != nil {
		log.Print("Ошибка при создании тикета: ", err)
		return err
	}
	return nil
}
