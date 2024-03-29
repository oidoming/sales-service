package client

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/services/clientsrv"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ClientHandler struct {
	ClientService clientsrv.Service
}

func NewClientHandler(client clientsrv.Service) *ClientHandler {
	return &ClientHandler{ClientService: client}
}

func (ch *ClientHandler) CreateClient(c echo.Context) error {
	client := domain.Client{}

	if err := c.Bind(&client); err != nil {
		message := domain.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	err := ch.ClientService.CreateClient(&client)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = client

	return c.JSON(http.StatusCreated, domain.MessageOK)
}

func (ch *ClientHandler) GetClients(c echo.Context) error {
	clients, err := ch.ClientService.GetClients()
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = clients

	return c.JSON(http.StatusOK, domain.MessageOK)
}

func (ch *ClientHandler) GetClient(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	client, err := ch.ClientService.GetClient(id)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = client

	return c.JSON(http.StatusOK, domain.MessageOK)
}

func (ch *ClientHandler) UpdateClient(c echo.Context) error {
	client := domain.Client{}
	if err := c.Bind(&client); err != nil {
		message := domain.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	id, _ := uuid.Parse(c.Param("id"))

	err := ch.ClientService.UpdateClient(id, client)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = client

	return c.JSON(http.StatusOK, domain.MessageOK)
}

func (ch *ClientHandler) DeleteClient(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	err := ch.ClientService.DeleteClient(id)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	m := domain.Message{
		Success: true,
	}

	return c.JSON(http.StatusOK, m)
}
