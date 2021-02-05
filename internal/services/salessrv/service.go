package salessrv

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	NewSale(saleRequest domain.SaleRequest) error
	GetSales() (domain.SalesResponse, error)
	GetSale(sale_id string) (domain.SaleResponse, error)
}

type Repository interface {
	GetProduct(id string) (domain.Product, error)
	GetSales() (domain.SalesResponse, error)
	GetSale(sale_id string) (domain.SaleResponse, error)
	NewSaleConcept(concept domain.Concept) error
	NewSaleTransaction(sale domain.Sale) (string, error)
}

type salesService struct {
	r Repository
}

func NewService(r Repository) Service {
	return &salesService{r: r}
}


func (s *salesService) GetSales() (domain.SalesResponse, error) {
	sales, err := s.r.GetSales()

	return sales, err
}

func (s *salesService) GetSale(sale_id string) (domain.SaleResponse, error) {
	sale, err := s.r.GetSale(sale_id)
	if err != nil {
		return sale, err
	}

	return sale, nil
}

func (s *salesService) NewSale(saleRequest domain.SaleRequest) error {
	product, err := s.r.GetProduct(saleRequest.ProductID)
	if err != nil {
		return err
	}

	total := calcTotal(saleRequest.Quantity, product.UnitPrice)

	sale := domain.Sale{
		ID:       uuid.New().String(),
		ClientID: saleRequest.ClientID,
		Total:    total,
	}

	saleID, err := s.r.NewSaleTransaction(sale)
	if err != nil {
		return err
	}

	concept := domain.Concept{
		ID:        uuid.New().String(),
		SaleID:    saleID,
		ProductID: product.ID,
		Quantity:  saleRequest.Quantity,
		UnitPrice: product.UnitPrice,
		Total:     total,
	}

	err = s.r.NewSaleConcept(concept)
	if err != nil {
		return err
	}

	return nil
}

func calcTotal(quantity int, unitPrice float64) float64 {
	return float64(quantity) * unitPrice
}
