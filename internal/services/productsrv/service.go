package productsrv

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	AddProduct(product *domain.Product) error
	GetProduct(id string) (domain.Product, error)
	GetProducts() (domain.Products, error)
	UpdateProduct(id string, product domain.Product) error
	RemoveProduct(id string) error
}

type Repository interface {
	AddProduct(product *domain.Product) error
	GetProduct(id string) (domain.Product, error)
	GetProducts() (domain.Products, error)
	UpdateProduct(id string, product domain.Product) error
	RemoveProduct(id string) error
}

type productsService struct {
	r Repository
}

func NewService(r Repository) Service {
	return &productsService{r: r}
}

func (s *productsService) AddProduct(product *domain.Product) error {
	product.ID = uuid.New().String()
	err := s.r.AddProduct(product)

	return err
}

func (s *productsService) GetProduct(id string) (domain.Product, error) {
	product, err := s.r.GetProduct(id)

	return product, err
}

func (s *productsService) GetProducts() (domain.Products, error) {
	products, err := s.r.GetProducts()

	return products, err
}

func (s *productsService) UpdateProduct(id string, product domain.Product) error {
	err := s.r.UpdateProduct(id, product)

	return err
}

func (s *productsService) RemoveProduct(id string) error {
	err := s.r.RemoveProduct(id)

	return err
}
