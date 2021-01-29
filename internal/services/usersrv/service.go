package usersrv

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/services"
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

	hash, err := services.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	err = s.repository.InsertUser(user)
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
	u, err := s.GetUser(id)
	if err != nil {
		return err
	}

	if user.Name == "" {
		user.Name = u.Name
	}

	if user.Email == "" {
		user.Email = u.Email
	}

	if user.Password != "" {
		hash, err := services.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = hash
	} else {
		user.Password = u.Password
	}

	err = s.repository.UpdateUser(id, user)
	if err != nil {
		return err
	}

	return nil
}
