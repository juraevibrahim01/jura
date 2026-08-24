package repository

import (
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
