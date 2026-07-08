package service

import (
	"log"

	"github.com/juraevibrahim01/jura/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type User_service struct {
	repository *repository.User_repository
}

func User_new_service(repository *repository.User_repository) *User_service {
	return &User_service{repository: repository}
}

func (u *User_service) User_create(email, password *string) error {

	// Проверка на сушествования пользователя
	err := u.User_check_exist_user(email)
	if err != nil {
		return err
	}

	// Хещирования пороля
	hash_password, err := u.User_generate_hash_password(password)
	if err != nil {
		return err
	}

	err = u.repository.User_create(email, &hash_password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User_service) User_generate_hash_password(password *string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), 12)
	if err != nil {
		log.Println("Ошибка при хешировании пороля!")
		return "", err
	}
	return string(hash), nil
}

func (u *User_service) User_check_exist_user(email *string) error {
	err := u.repository.User_check_exist_user(email)
	if err != nil {
		return err
	}
	return nil
}

func (u *User_service) GetUserRole(email string) (string, error) {
	return u.repository.GetUserRole(email)
}
