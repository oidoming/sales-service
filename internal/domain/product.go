package domain

type Product struct {
	tableName struct{} `pg:"product"`
	ID string `json:"id" pg:"id"`
	Name string `json:"name" pg:"name"`
	UnitPrice float64 `json:"unit_price" pg:"unit_price"`
	Cost float64 `json:"cost" pg:"cost"`
}

type Products []Product