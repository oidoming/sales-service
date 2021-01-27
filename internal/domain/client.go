package domain

import "github.com/google/uuid"

type Client struct {
	tableName struct{} `pg:"client"`
	ID        uuid.UUID      `json:"id" pg:"id,pk"`
	Name      string   `json:"name" pg:"name"`
}

type Clients []Client
