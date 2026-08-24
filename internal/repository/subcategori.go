package repository

import (
	"database/sql"
	"log"

	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/pkg"
)

type SubCategories struct {
	postgres *pkg.Postgres
}

func NewSubCategories(postgres *pkg.Postgres) *SubCategories {
	return &SubCategories{postgres: postgres}
}

func (s *SubCategories) GetSubCategories(categoriID *int) (*[]models.SubCategoriResdb, error) {

	var subCategories []models.SubCategoriResdb
	var subCategori models.SubCategoriResdb

	query := `
		SELECT id, "name"
		FROM subcategories
		WHERE categori_id = $1
	`
	rows, err := s.postgres.DB.Query(query, categoriID)
	if err != nil {
		log.Println("Ошибка при получении подкатегрий: ", err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&subCategori.ID, &subCategori.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				return &[]models.SubCategoriResdb{}, err
			}
			log.Println("Ошибка при сканоровании подкатегорий: ", err)
			return nil, err
		}
		subCategories = append(subCategories, subCategori)
	}

	if rows.Err(); err != nil {
		log.Println("Ошибка при итерации по строкам подкатегорий", err)
		return nil, err
	}

	return &subCategories, nil

}
