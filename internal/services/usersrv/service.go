package usersrv

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	CreateUser(user *domain.User) error
	GetUser(id uuid.UUID) (domain.User, error)
	UpdateUser(id uuid.UUID, user domain.User) error
}

type Repository interface {
	InsertUser(user *domain.User) error
	SelectUser(id uuid.UUID) (domain.User, error)
	UpdateUser(id uuid.UUID, user domain.User) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateUser(user *domain.User) error {
	user.ID = uuid.New()
	err := s.repository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUser(id uuid.UUID) (domain.User, error) {
	user, err := s.repository.SelectUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateUser(id uuid.UUID, user domain.User) error {
	err := s.repository.UpdateUser(id, user)
	if err != nil {
		return err
	}

	return nil
}
