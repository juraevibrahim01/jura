package repository

import (
	"database/sql"
	"log"

	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/pkg"
)

type CategoriReposotori struct {
	postgres *pkg.Postgres
}

func NewReCategori(postgres *pkg.Postgres) *CategoriReposotori {
	return &CategoriReposotori{postgres: postgres}
}

func (r *CategoriReposotori) GetCategoti(projectID *int) (*[]models.Categoridbres, error) {

	var categories []models.Categoridbres

	query := `
		SELECT id, "name"
		FROM categories
		WHERE project_id = $1
		`
	rows, err := r.postgres.DB.Query(query, projectID)
	if err != nil {
		log.Print("Ошибка при получении категориев: , ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var categori models.Categoridbres

		err = rows.Scan(&categori.ID, &categori.Name)
		if err != nil {
			log.Println("Ошибка при сканировании категориев: ", err)
			return nil, err
		}

		categories = append(categories, categori)
	}

	if rows.Err() != nil {
		log.Print("Ошибка при итерации по строкам категорийы: ", rows.Err())
		return nil, rows.Err()
	}

	return &categories, nil

}

func (r *CategoriReposotori) CreateCategory(name *string, projectID *int) (string, error) {

	query := `
		INSERT INTO categories ("name", project_id)
		VALUES ($1, $2)
	`

	_, err := r.postgres.DB.Exec(query, name, projectID)
	if err != nil {
		log.Print("Ошибка при создании категории: ", err)
		return "", err
	}

	return "success", nil
}

func (r *CategoriReposotori) DeleteCategory(category_id *int) (string, error) {
	tx, err := r.postgres.DB.Begin()
	if err != nil {
		log.Print("Ошибка при старте транзакции: ", err)
		return "", err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(
		`DELETE FROM subcategories WHERE categori_id = $1`, *category_id,
	); err != nil {
		log.Print("Ошибка при удалении подкатегорий: ", err)
		return "", err
	}

	res, err := tx.Exec(
		`DELETE FROM categories WHERE id = $1`, *category_id,
	)
	if err != nil {
		log.Print("Ошибка при удалении категории: ", err)
		return "", err
	}

	if n, _ := res.RowsAffected(); n == 0 {
		return "", sql.ErrNoRows // категории с таким id нет
	}

	if err := tx.Commit(); err != nil {
		log.Print("Ошибка при коммите: ", err)
		return "", err
	}

	return "success", nil

}

func (r *CategoriReposotori) CheckSubcategoriesExist(categoryID *int) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM subcategories WHERE categori_id = $1
		)
	`

	var exists bool
	if err := r.postgres.DB.QueryRow(query, categoryID).Scan(&exists); err != nil {
		log.Printf("Ошибка при проверке подкатегорий (category_id=%d): %v", categoryID, err)
		return false, err
	}

	log.Printf("category_id=%d, подкатегории найдены: %v", categoryID, exists)
	return exists, nil
}
