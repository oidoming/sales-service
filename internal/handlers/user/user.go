package user

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/handlers/response"
	"github.com/Oscar-inc117/sales-service/internal/services/usersrv"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUser(service usersrv.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := domain.User{}

		if err := c.Bind(&user); err != nil {
			message := response.CreateErrorResponse(http.StatusBadRequest, err)
			return c.JSON(http.StatusBadRequest, message)
		}

		err := service.CreateUser(&user)
		if err != nil {
			message := response.CreateErrorResponse(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, message)
		}

		response.MessageOK.Payload = user

		return c.JSON(http.StatusCreated, response.MessageOK)
	}
}

func GetUser(service usersrv.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		id, _ := uuid.Parse(c.Param("id"))

		user, err := service.GetUser(id)
		if err != nil {
			message := response.CreateErrorResponse(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, message)
		}

		response.MessageOK.Payload = user

		return c.JSON(http.StatusOK, response.MessageOK)
	}
}

func UpdateUser(service usersrv.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := domain.User{}
		if err := c.Bind(&user); err != nil {
			message := response.CreateErrorResponse(http.StatusBadRequest, err)
			return c.JSON(http.StatusBadRequest, message)
		}

		id, _ := uuid.Parse(c.Param("id"))

		err := service.UpdateUser(id, user)
		if err != nil {
			message := response.CreateErrorResponse(http.StatusInternalServerError, err)
			return c.JSON(http.StatusInternalServerError, message)
		}

		response.MessageOK.Payload = user

		return c.JSON(http.StatusOK, response.MessageOK)
	}
}
