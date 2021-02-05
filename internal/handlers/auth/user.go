package auth

import (
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *AuthHandler) CreateUser(c echo.Context) error {
	user := domain.User{}

	if err := c.Bind(&user); err != nil {
		message := domain.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	err := a.UserService.CreateUser(&user)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = user

	return c.JSON(http.StatusCreated, domain.MessageOK)
}

func (a *AuthHandler) GetUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	user, err := a.UserService.GetUser(id)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = user

	return c.JSON(http.StatusOK, domain.MessageOK)
}

func (a *AuthHandler) UpdateUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		message := domain.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	err := a.UserService.UpdateUser(id, user)
	if err != nil {
		message := domain.CreateErrorResponse(http.StatusInternalServerError, err)
		return c.JSON(http.StatusInternalServerError, message)
	}

	domain.MessageOK.Payload = user

	return c.JSON(http.StatusOK, domain.MessageOK)
}
