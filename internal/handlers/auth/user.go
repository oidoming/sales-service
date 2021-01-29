package auth

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/handlers/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *AuthHandler) CreateUser(c echo.Context) error {
	user := domain.User{}

	if err := c.Bind(&user); err != nil {
		message := response.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	err := a.UserService.CreateUser(&user)
	if err != nil {
		message := response.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	response.MessageOK.Payload = user

	return c.JSON(http.StatusCreated, response.MessageOK)
}

func (a *AuthHandler) GetUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	user, err := a.UserService.GetUser(id)
	if err != nil {
		message := response.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	response.MessageOK.Payload = user

	return c.JSON(http.StatusOK, response.MessageOK)
}

func (a *AuthHandler) UpdateUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		message := response.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	err := a.UserService.UpdateUser(id, user)
	if err != nil {
		message := response.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	response.MessageOK.Payload = user

	return c.JSON(http.StatusOK, response.MessageOK)
}
