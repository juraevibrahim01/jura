package service

import "github.com/juraevibrahim01/jura/internal/repository"

type PopularService interface {
	GetPopular() repository.PopularResult
}

type RecentService interface {
	GetPopular() repository.PopularResult
}

type BuyAgainService interface {
	GetPopular() repository.PopularResult
}

type popularService struct {
	repo repository.PopularRepository
}

type recentService struct {
	repo repository.PopularRepository
}

type buyAgainService struct {
	repo repository.PopularRepository
}

func NewPopularService(repo repository.PopularRepository) PopularService {
	return &popularService{
		repo: repo,
	}
}

func NewRecentService(repo repository.PopularRepository) RecentService {
	return &recentService{
		repo: repo,
	}
}

func NewBuyAgainService(repo repository.PopularRepository) BuyAgainService {
	return &buyAgainService{
		repo: repo,
	}
}

func (s *popularService) GetPopular() repository.PopularResult {
	return s.repo.GetPopular()
}

func (s *recentService) GetPopular() repository.PopularResult {
	return s.repo.GetPopular()
}

func (s *buyAgainService) GetPopular() repository.PopularResult {
	return s.repo.GetPopular()
}
