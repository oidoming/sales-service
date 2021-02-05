package sales

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/services/salessrv"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SalesHandler struct {
	SalesService salessrv.Service
}

func NewSalesHandler(sales salessrv.Service) *SalesHandler {
	return &SalesHandler{
		SalesService: sales,
	}
}

func (s *SalesHandler) NewSale(c echo.Context) error {
	sale := domain.SaleRequest{}

	if err := c.Bind(&sale); err != nil {
		message := domain.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	if err := s.SalesService.NewSale(sale); err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = sale

	return c.JSON(http.StatusCreated, domain.MessageOK)
}

func (s *SalesHandler) GetSale(c echo.Context) error {
	id := c.Param("id")

	sale, err := s.SalesService.GetSale(id)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = sale

	return c.JSON(http.StatusCreated, domain.MessageOK)
}

func (s *SalesHandler) GetSales(c echo.Context) error {
	sales, err := s.SalesService.GetSales()
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = sales

	return c.JSON(http.StatusCreated, domain.MessageOK)
}
