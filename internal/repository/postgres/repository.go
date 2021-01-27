package postgres

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type ClientStorage struct {
	DB *pg.DB
}

func (s *ClientStorage) InsertClient(client *domain.Client) error {

	if _, err := s.DB.Model(client).Insert(); err != nil {
		return err
	}

	return nil
}

func (s *ClientStorage) SelectClients() (domain.Clients, error) {
	var clients domain.Clients

	err := s.DB.Model(&clients).Select()

	return clients, err
}

func (s *ClientStorage) SelectClient(id uuid.UUID) (domain.Client, error) {
	var client domain.Client

	err := s.DB.Model(&client).Where("id=?", id).Select()

	return client, err
}

func (s *ClientStorage) UpdateClient(id uuid.UUID, client domain.Client) error {
	_, err := s.DB.Model(&client).Where("id=?", id).Update()

	return err
}

func (s *ClientStorage) DeleteClient(id uuid.UUID) error {
	var client domain.Client

	_, err := s.DB.Model(&client).Where("id=?", id).Delete()

	return err
}