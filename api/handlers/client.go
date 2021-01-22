package handlers

import (
	"net/http"
	"strconv"

	"github.com/Oscar-inc117/sales-service/internal/models"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
)

type Client struct {
	DB *pg.DB
}

func (cl *Client) CreateClient(c echo.Context) error {
	client := &models.Client{}

	if err := c.Bind(client); err != nil {
		return err
	}

	err := client.InsertClient(cl.DB)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, client)
}

func (cl *Client) GetClients(c echo.Context) error {
	client := &models.Client{}
	clients, err := client.SelectClients(cl.DB)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, clients)
}

func (cl *Client) GetClient(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	client := &models.Client{}
	client, err := client.SelectClient(cl.DB, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, client)
}

func (cl *Client) UpdateClient(c echo.Context) error {
	client := &models.Client{}
	if err := c.Bind(client); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err := client.UpdateClient(cl.DB, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, client)
}

func (cl *Client) DeleteClient(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	client := &models.Client{}
	err := client.DeleteClient(cl.DB, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, client)
}
