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

func (r *Test_keys_repository) GetTestKeys(project_id, category_id, subcategoryID *int) ([]models.TestKey, error) {
	query := `
		SELECT
			tk.id,
			tk.date,
			tk.name,
			tk.module,
			tk.precondition,
			tk.steps,
			tk.expectation_res,
			tk.actual_res,
			tk.comment
		FROM test_keys tk
		JOIN subcategories sc on sc.id = tk.subcategory_id
		JOIN categories c
			on c.id = sc.categori_id
		JOIN projects p
			on p.id = c.project_id
		where p.id = $1 and c.id = $2 and sc.id = $3;
	`

	rows, err := r.postgres.DB.Query(query, *project_id, *category_id, *subcategoryID)
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

	if rows.Err() != nil {
		log.Println("Ошибка при итерации по строкам тестовых ключей: ", rows.Err())
		return nil, rows.Err()
	}

	return testKeys, nil
}

func (r *Test_keys_repository) GetTestKeyByID(id *int, user_id *int) (*models.TestKey, error) {
	query := `
		SELECT
			tk.id,
			tk.date,
			tk.name,
			tk.module,
			tk.precondition,
			tk.steps,
			tk.expectation_res,
			tk.actual_res,
			tk.comment
		FROM test_keys tk
		JOIN subcategories sc on sc.id = tk.subcategory_id
		JOIN categories c on c.id = sc.categori_id
		JOIN projects p on p.id = c.project_id
		JOIN users u on u.id = p.user_id
		WHERE tk.id = $1 and u.id = $2;
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

func (r *Test_keys_repository) CreateTestKey(request *models.TestKeyCreateRequest, subcategory_id *int) error {
	query := `
		INSERT INTO test_keys (date, name, module, precondition, steps, expectation_res, actual_res, comment, subcategory_id, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
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
		*subcategory_id,
		request.Status,
	)
	if err != nil {
		log.Print("Ошибка при создании тестового кейса: ", err)
		return err
	}

	return nil
}

func (r *Test_keys_repository) GetCategories() ([]models.Category, error) {
	query := `
		SELECT id, name
		FROM projects;
	`
	rows, err := r.postgres.DB.Query(query)
	if err != nil {
		log.Print("Ошибка при получении категорий: ", err)
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			log.Print("Ошибка при сканировании категорий: ", err)
			return nil, err
		}
		categories = append(categories, category)
	}

	if rows.Err() != nil {
		log.Println("Ошибка при итерации по строкам категорий: ", rows.Err())
		return nil, rows.Err()
	}

	return categories, nil
}

func (r *Test_keys_repository) GetProjects() ([]models.Project, error) {
	query := `
		SELECT id, name
		FROM projects;
	`

	rows, err := r.postgres.DB.Query(query)
	if err != nil {
		log.Print("Ошибка при получении проектов: ", err)
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err = rows.Scan(&p.ID, &p.Name)
		if err != nil {
			log.Print("Ошибка при сканировании проектов: ", err)
			return nil, err
		}
		projects = append(projects, p)
	}

	if rows.Err() != nil {
		log.Println("Ошибка при итерации по строкам проектов: ", rows.Err())
		return nil, rows.Err()
	}

	return projects, nil
}

func (r *Test_keys_repository) GetProjectByID(id int) (*models.ProjectID, error) {
	query := `
		select p.name as project_name, COUNT(DISTINCT t.ID) as testkeys_total, COUNT(DISTINCT b.ID) as tickets_total
		from test_keys t
		join tickets b on b.project_id = t.project_id
		join projects p on p.id = t.project_id
		where p.id = $1
		group by p.name
	`

	row := r.postgres.DB.QueryRow(query, id)
	var p models.ProjectID

	if err := row.Scan(&p.ProjectName, &p.TestKeys_total, &p.Tickets_total); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Print("Ошибка при получении проекта по id: ", err)
		return nil, err
	}

	return &p, nil
}
