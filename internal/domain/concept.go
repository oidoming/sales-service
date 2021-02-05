package domain

type Concept struct {
	tableName struct{} `pg:"concept"`
	ID        string   `json:"id" pg:"id"`
	SaleID    string   `json:"sale_id" pg:"sale_id"`
	ProductID string   `json:"product_id" pg:"product_id"`
	Quantity  int      `json:"quantity" pg:"quantity"`
	UnitPrice float64  `json:"unit_price" pg:"unit_price"`
	Total     float64  `json:"total" pg:"total"`
}