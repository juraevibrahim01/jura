package repository

import (
	"database/sql"
	"log"

	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/pkg"
)

type User_repository struct {
	postgres *pkg.Postgres
}

func User_new_repository(postgres *pkg.Postgres) *User_repository {
	return &User_repository{postgres: postgres}
}

func (u *User_repository) User_create(email, password *string) error {
	query := `
	insert into users(email, password)
	values($1, $2)
	`
	_, err := u.postgres.DB.Exec(query, email, password)
	if err != nil {
		log.Print("Ошибка при insert данных: ", err)
		return err
	}
	return nil
}

func (u *User_repository) User_check_exist_user(email *string) error {
	var res_db string

	query := `
		select email
		from users
		where email = $1;
	`
	row := u.postgres.DB.QueryRow(query, email)
	err := row.Scan(&res_db)
	if err == sql.ErrNoRows {
		log.Print("Ошибка: Запрос не нашел ни одно пользователя с почтой: ", *email)
		return nil
	}
	if err != nil {
		log.Print("Ошибка при сканировании", err)
		return err
	}
	return models.User_err_exists_user
}

func (u *User_repository) GetUserRole(email string) (string, error) {
	var role string

	query := `
		SELECT r.name
		FROM roles r
		JOIN user_roles ur on r.id = ur.roles_id
		JOIN users u on u.id = ur.user_id
        WHERE u.email = $1;
	`
	row := u.postgres.DB.QueryRow(query, email)
	err := row.Scan(&role)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		log.Print("Ошибка при получении роли пользователя: ", err)
		return "", err
	}

	return role, nil
}
