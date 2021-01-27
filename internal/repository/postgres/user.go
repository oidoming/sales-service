package postgres

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) InsertUser(user *domain.User) error {

	if _, err := u.DB.Model(user).Insert(); err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) SelectUser(id uuid.UUID) (domain.User, error) {
	var user domain.User

	err := u.DB.Model(&user).Where("id=?", id).Select()

	return user, err
}

func (u *UserRepo) UpdateUser(id uuid.UUID, user domain.User) error {
	_, err := u.DB.Model(&user).Where("id=?", id).Update()

	return err
}
