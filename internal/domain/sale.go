package domain

import "time"

type Sale struct {
	tableName struct{} `pg:"sale"`
	ID string `json:"id" pg:"id"`
	ClientID string `json:"client_id" pg:"client_id"`
	Date time.Time `json:"date" pg:"date"`
	Total float64 `json:"total" pg:"total"`
}

type SaleResponse struct {
	Date        string    `json:"date"`
	ClientName  string    `json:"client_name"`
	ProductName string    `json:"product_name"`
	UnitPrice   float64   `json:"unit_price"`
	Quantity    int       `json:"quantity"`
	Total       float64   `json:"total"`
}

type SalesResponse []SaleResponse

type SaleRequest struct {
	ClientID    string    `json:"client_id"`
	ProductID   string    `json:"product_id"`
	UnitPrice   float64   `json:"unit_price"`
	Quantity    int       `json:"quantity"`
	Total       float64   `json:"total"`
}