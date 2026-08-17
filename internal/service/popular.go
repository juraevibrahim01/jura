package service

import "github.com/juraevibrahim01/jura/internal/repository"

type PopularService interface {
	GetPopular() repository.PopularResult
	GetPopularNull() repository.PopularResult
}

type popularService struct {
	repo repository.PopularRepository
}

func NewPopularService(repo repository.PopularRepository) PopularService {
	return &popularService{
		repo: repo,
	}
}

func (s *popularService) GetPopular() repository.PopularResult {
	return s.repo.GetPopular()
}

func (s *popularService) GetPopularNull() repository.PopularResult {
	return s.repo.GetPopularNull()
}