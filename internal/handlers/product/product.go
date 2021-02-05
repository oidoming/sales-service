package product

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/services/productsrv"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductHandler struct {
	ProductService productsrv.Service
}

func NewProductHandler(product productsrv.Service) *ProductHandler {
	return &ProductHandler{
		ProductService: product,
	}
}

func (p *ProductHandler) AddProduct(c echo.Context) error {
	product := new(domain.Product)

	if err := c.Bind(product); err != nil {
		message := domain.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	if err := p.ProductService.AddProduct(product); err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = product

	return c.JSON(http.StatusCreated, domain.MessageOK)
}

func (p *ProductHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")

	product, err := p.ProductService.GetProduct(id)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = product

	return c.JSON(http.StatusOK, domain.MessageOK)
}

func (p *ProductHandler) GetProducts(c echo.Context) error {
	products, err := p.ProductService.GetProducts()
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = products

	return c.JSON(http.StatusOK, domain.MessageOK)
}

func (p *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	product := domain.Product{}

	if err := c.Bind(&product); err != nil {
		message := domain.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	if err := p.ProductService.UpdateProduct(id, product);err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = product

	return c.JSON(http.StatusOK, domain.MessageOK)
}

func (p *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	if err := p.ProductService.RemoveProduct(id); err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	m := domain.Message{
		Success: true,
	}

	return c.JSON(http.StatusOK, m)
}
