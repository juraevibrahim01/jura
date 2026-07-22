package service

import (
	"github.com/juraevibrahim01/jura/internal/models"
	"github.com/juraevibrahim01/jura/internal/repository"
)

type Test_keys_service struct {
	repository *repository.Test_keys_repository
}

func New_Test_keys_service(repository *repository.Test_keys_repository) *Test_keys_service {
	return &Test_keys_service{repository: repository}
}

func (s *Test_keys_service) GetTestKeys(user_id *int) ([]models.TestKey, error) {
	return s.repository.GetTestKeys(user_id)
}

func (s *Test_keys_service) GetTestKeyByID(id *int, user_id *int) (*models.TestKey, error) {
	return s.repository.GetTestKeyByID(id, user_id)
}

func (s *Test_keys_service) CreateTestKey(request *models.TestKeyCreateRequest, user_id *int) error {
	return s.repository.CreateTestKey(request, user_id)
}
