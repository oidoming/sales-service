package database

import (
	"context"

	"github.com/Oscar-inc117/sales-service/internal/config"
	"github.com/go-pg/pg/v10"
)

func GetConnection(c *config.DatabaseConfig) *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     c.User,
		Password: c.Password,
		Database: c.DBName,
	})

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	return db
}
