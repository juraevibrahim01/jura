package service

import "github.com/juraevibrahim01/jura/internal/repository"

type TestService interface {
	GetTest() repository.TestResult
}

type testService struct {
	repo repository.TestRepository
}

func NewForYouService(repo repository.TestRepository) TestService {
	return &testService{
		repo: repo,
	}
}

func (s *testService) GetTest() repository.TestResult {
	return s.repo.GetTest()
}
