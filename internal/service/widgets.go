package service

import "github.com/juraevibrahim01/jura/internal/repository"

type WidgetsService interface {
	GetWidgets() repository.WidgetsResult
}

type widgetsService struct {
	repo repository.WidgetsRepository
}

func NewWidgetsService(repo repository.WidgetsRepository) WidgetsService {
	return &widgetsService{repo: repo}
}

func (s *widgetsService) GetWidgets() repository.WidgetsResult {
	return s.repo.GetWidgets()
}
