package domain

import "github.com/google/uuid"

type User struct {
	tableName struct{} `pg:"user_account"`
	ID uuid.UUID `json:"id" pg:"id"`
	Name string `json:"name" pg:"name"`
	Email string `json:"email" pg:"email"`
	Password string `json:"password" pg:"password"`
}
