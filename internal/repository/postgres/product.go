package postgres

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
)


func (r *Repository) AddProduct(product *domain.Product) error {
	_, err := r.DB.Model(product).Insert()

	return err
}

func (r *Repository) GetProduct(id string) (domain.Product, error) {
	var product domain.Product

	err := r.DB.Model(&product).Where("id=?", id).Select()

	return product, err
}

func (r *Repository) GetProducts() (domain.Products, error) {
	var products domain.Products

	err := r.DB.Model(&products).Select()

	return products, err
}

func (r *Repository) UpdateProduct(id string, product domain.Product) error {
	_, err := r.DB.Model(&product).Where("id=?", id).Update()

	return err
}

func (r *Repository) RemoveProduct(id string) error {
	var product domain.Product

	_, err := r.DB.Model(&product).Where("id=?", id).Delete()

	return err
}
