package client

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/handlers/response"
	"github.com/Oscar-inc117/sales-service/internal/services/clientsrv"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)


func CreateClient(service clientsrv.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		client := domain.Client{}

		if err := c.Bind(&client); err != nil {
			message := response.CreateErrorResponse(http.StatusBadRequest, err)
			return c.JSON(http.StatusBadRequest, message)
		}

		err := service.CreateClient(&client)
		if err != nil {
			message := response.CreateErrorResponse(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, message)
		}

		response.MessageOK.Payload = client

		return c.JSON(http.StatusCreated, response.MessageOK)
	}
}

func GetClients(service clientsrv.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		clients, err := service.GetClients()
		if err != nil {
			message := response.CreateErrorResponse(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, message)
		}

		response.MessageOK.Payload = clients

		return c.JSON(http.StatusOK, response.MessageOK)
	}
}

func GetClient(service clientsrv.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, _ := uuid.Parse(c.Param("id"))

		client, err := service.GetClient(id)
		if err != nil {
			message := response.CreateErrorResponse(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, message)
		}

		response.MessageOK.Payload = client

		return c.JSON(http.StatusOK, response.MessageOK)
	}
}

func UpdateClient(service clientsrv.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		client := domain.Client{}
		if err := c.Bind(&client); err != nil {
			message := response.CreateErrorResponse(http.StatusBadRequest, err)
			return c.JSON(http.StatusBadRequest, message)
		}

		id, _ := uuid.Parse(c.Param("id"))

		err := service.UpdateClient(id, client)
		if err != nil {
			message := response.CreateErrorResponse(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, message)
		}

		response.MessageOK.Payload = client

		return c.JSON(http.StatusOK, response.MessageOK)
	}
}

func DeleteClient(service clientsrv.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, _ := uuid.Parse(c.Param("id"))

		err := service.DeleteClient(id)
		if err != nil {
			message := response.CreateErrorResponse(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, message)
		}

		m := response.Message{
			Success: true,
		}

		return c.JSON(http.StatusOK, m)
	}
}
