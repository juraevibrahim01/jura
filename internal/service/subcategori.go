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

func (s *SubCategoriesService) GetSubCategories(categoriID *int64) (*[]models.SubCategoriResdb, error) {
	return s.repo.GetSubCategories(categoriID)
}

func (s *SubCategoriesService) CreateSubCategory(name *string, categoriID *int64) (string, error) {
	return s.repo.CreateSubCategory(name, categoriID)
}
