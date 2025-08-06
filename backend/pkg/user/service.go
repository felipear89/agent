package user

import (
	"errors"
	"fmt"
)

type Service struct {
	repo Repository
}

func newService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return nil, fmt.Errorf("not implemented")
	//return s.repo.FindAll()
}

func (s *Service) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) CreateUser(user User) (*User, error) {
	return s.repo.Create(user)
}

func (s *Service) UpdateUser(user User) (*User, error) {
	return s.repo.Update(user)
}

func (s *Service) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) Authenticate(email, password string) (*User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
