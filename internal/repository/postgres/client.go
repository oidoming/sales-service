package postgres

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/google/uuid"
)

func (r *Repository) InsertClient(client *domain.Client) error {

	if _, err := r.DB.Model(client).Insert(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) SelectClients() (domain.Clients, error) {
	var clients domain.Clients

	err := r.DB.Model(&clients).Select()

	return clients, err
}

func (r *Repository) SelectClient(id uuid.UUID) (domain.Client, error) {
	var client domain.Client

	err := r.DB.Model(&client).Where("id=?", id).Select()

	return client, err
}

func (r *Repository) UpdateClient(id uuid.UUID, client domain.Client) error {
	_, err := r.DB.Model(&client).Where("id=?", id).Update()

	return err
}

func (r *Repository) DeleteClient(id uuid.UUID) error {
	var client domain.Client

	_, err := r.DB.Model(&client).Where("id=?", id).Delete()

	return err
}
