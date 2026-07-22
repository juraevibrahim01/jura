package repository

import (
	"database/sql"
	"log"

	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/pkg"
)

type Test_keys_repository struct {
	postgres *pkg.Postgres
}

func New_Test_keys_repository(postgres *pkg.Postgres) *Test_keys_repository {
	return &Test_keys_repository{postgres: postgres}
}

func (r *Test_keys_repository) GetTestKeys(user_id *int) ([]models.TestKey, error) {
	query := `
		SELECT id, date, name, module, precondition, steps, expectation_res, actual_res, comment
		FROM test_keys
		WHERE user_id = $1;
	`

	rows, err := r.postgres.DB.Query(query, *user_id)
	if err != nil {
		log.Print("Ошибка при получении тестовых ключей: ", err)
		return nil, err
	}
	defer rows.Close()

	var testKeys []models.TestKey
	for rows.Next() {
		var testKey models.TestKey
		err = rows.Scan(
			&testKey.ID,
			&testKey.Date,
			&testKey.Name,
			&testKey.Module,
			&testKey.Precondition,
			&testKey.Steps,
			&testKey.ExpectationRes,
			&testKey.ActualRes,
			&testKey.Comment,
		)
		if err != nil {
			log.Print("Ошибка при сканировании тестовых ключей: ", err)
			return nil, err
		}
		testKeys = append(testKeys, testKey)
	}

	return testKeys, nil
}

func (r *Test_keys_repository) GetTestKeyByID(id *int, user_id *int) (*models.TestKey, error) {
	query := `
		SELECT id, date, name, module, precondition, steps, expectation_res, actual_res, comment
		FROM test_keys
		WHERE id = $1 AND user_id = $2;
	`

	row := r.postgres.DB.QueryRow(query, id, user_id)
	var testKey models.TestKey
	if err := row.Scan(
		&testKey.ID,
		&testKey.Date,
		&testKey.Name,
		&testKey.Module,
		&testKey.Precondition,
		&testKey.Steps,
		&testKey.ExpectationRes,
		&testKey.ActualRes,
		&testKey.Comment,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Print("Ошибка при получении тестового кейса по id: ", err)
		return nil, err
	}

	return &testKey, nil
}

func (r *Test_keys_repository) CreateTestKey(request *models.TestKeyCreateRequest, user_id *int) error {
	query := `
		INSERT INTO test_keys (date, name, module, precondition, steps, expectation_res, actual_res, comment, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := r.postgres.DB.Exec(query,
		request.Date,
		request.Name,
		request.Module,
		request.Precondition,
		request.Steps,
		request.ExpectationRes,
		request.ActualRes,
		request.Comment,
		*user_id,
	)
	if err != nil {
		log.Print("Ошибка при создании тестового кейса: ", err)
		return err
	}

	return nil
}
