package postgres

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/google/uuid"
)


func (r *Repository) InsertUser(user *domain.User) error {

	if _, err := r.DB.Model(user).Insert(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) SelectUser(id uuid.UUID) (domain.User, error) {
	var user domain.User

	err := r.DB.Model(&user).Where("id=?", id).Select()

	return user, err
}

func (r *Repository) UpdateUser(id uuid.UUID, user domain.User) error {
	_, err := r.DB.Model(&user).Where("id=?", id).Update()

	return err
}

func (r *Repository) SelectUserByEmail(email string) (domain.User, error) {
	var user domain.User

	err := r.DB.Model(&user).Where("email=?", email).Select()

	return user, err
}
