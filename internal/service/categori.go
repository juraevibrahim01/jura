package service

import (
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/repository"
)

type CategotiService struct {
	repo *repository.CategoriReposotori
}

func NewCategori(repo *repository.CategoriReposotori) *CategotiService {
	return &CategotiService{repo: repo}
}

func (c *CategotiService) GetCategories(projectID *int) (*[]models.Categoridbres, error) {
	return c.repo.GetCategoti(projectID)
}
