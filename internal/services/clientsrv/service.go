package clientsrv

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	CreateClient(client *domain.Client) error
	GetClients() (domain.Clients, error)
	GetClient(id uuid.UUID) (domain.Client, error)
	UpdateClient(id uuid.UUID, client domain.Client) error
	DeleteClient(id uuid.UUID) error
}

type Repository interface {
	InsertClient(client *domain.Client) error
	SelectClients() (domain.Clients, error)
	SelectClient(id uuid.UUID) (domain.Client, error)
	UpdateClient(id uuid.UUID, client domain.Client) error
	DeleteClient(id uuid.UUID) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateClient(client *domain.Client) error {
	client.ID = uuid.New()
	err := s.repository.InsertClient(client)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetClients() (domain.Clients, error) {
	clients, err := s.repository.SelectClients()
	if err != nil {
		return clients, err
	}

	return clients, nil
}

func (s *service) GetClient(id uuid.UUID) (domain.Client, error) {
	client, err := s.repository.SelectClient(id)
	if err != nil {
		return client, err
	}

	return client, nil
}

func (s *service) UpdateClient(id uuid.UUID, client domain.Client) error {
	err := s.repository.UpdateClient(id, client)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteClient(id uuid.UUID) error {
	err := s.repository.DeleteClient(id)
	if err != nil {
		return err
	}

	return nil
}