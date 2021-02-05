package postgres

import (
	"github.com/go-pg/pg/v10"
)


type Repository struct {
	DB *pg.DB
}

func NewRepository(DB *pg.DB) Repository {
	return Repository{DB: DB}
}
