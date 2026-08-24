package service

import (
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/repository"
)

type SubCategoriesService struct {
	repo *repository.SubCategories
}

func NewSubCategori(repo *repository.SubCategories) *SubCategoriesService {
	return &SubCategoriesService{repo: repo}
}

func (s *SubCategoriesService) GetSubCategories(categoriID *int) (*[]models.SubCategoriResdb, error) {
	return s.repo.GetSubCategories(categoriID)
}
