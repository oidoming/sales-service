package postgres

import (
	"fmt"
	"github.com/Oscar-inc117/sales-service/internal/domain"
)


func (r *Repository) GetSales() (domain.SalesResponse, error) {
	var sales domain.SalesResponse
	_, err := r.DB.Query(&sales,
		`select s.date, cl.name as client_name, p.name as product_name, c.unit_price, c.quantity, c.total
			   from concept c 
			   inner join sale s on c.sale_id = s.id
			   inner join client cl on s.client_id = cl.id
			   inner join product p on c.product_id = p.id`)

	if err != nil {
		return sales, err
	}

	return sales, nil
}

func (r *Repository) GetSale(sale_id string) (domain.SaleResponse, error) {
	var sale domain.SaleResponse
	_, err := r.DB.Query(&sale,
		`select s.date, cl.name as client_name, p.name as product_name, c.unit_price, c.quantity, c.total
			   from concept c 
			   inner join sale s on c.sale_id = s.id
			   inner join client cl on s.client_id = cl.id
			   inner join product p on c.product_id = p.id
			   where s.id = ?`, sale_id)

	if err != nil {
		return sale, err
	}

	return sale, nil
}

func (r *Repository) NewSaleConcept(concept domain.Concept) error {
	_, err := r.DB.Model(&concept).Insert()

	return err
}

func (r *Repository) NewSaleTransaction(sale domain.Sale) (string, error) {
	_, err := r.DB.Model(&sale).Returning("id").Insert()
	if err != nil {
		return "", err
	}
	fmt.Println(sale.ID)
	return sale.ID, err
}