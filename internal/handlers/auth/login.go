package auth

import (
	"errors"
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/handlers/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *AuthHandler) Login(c echo.Context) error {
	user := domain.User{}

	if err := c.Bind(&user); err != nil {
		message := response.CreateErrorResponse(http.StatusBadRequest, err)
		return c.JSON(http.StatusBadRequest, message)
	}

	user, ok := a.AuthService.Auth(user.Email, user.Password)
	if !ok {
		message := response.CreateErrorResponse(http.StatusUnauthorized, errors.New("incorrect email or password"))
		return c.JSON(http.StatusUnauthorized, message)
	}

	accessToken := a.AuthService.GenerateJWT(user)

	refreshToken := a.AuthService.GenerateRefreshJWT(user)

	resp := response.Message{
		Success: true,
		Payload: map[string]interface{}{
			"access_token": accessToken,
			"refreshToken": refreshToken,
			"auth": map[string]string{
				"name" : user.Name,
				"email": user.Email,
			},
		},
	}

	return c.JSON(http.StatusCreated, resp)
}

func (a *AuthHandler) RefreshToken(c echo.Context) error {
	user := domain.User{}

	userId := c.Get("user_id")

	uid, _ := uuid.Parse(userId.(string))

	user, err := a.UserService.GetUser(uid)
	if err != nil {
		return err
	}

	accessToken := a.AuthService.GenerateJWT(user)

	m := response.Message{
		Success: true,
		Payload: map[string]string{
			"accessToken": accessToken,
		},
	}

	return c.JSON(http.StatusOK, m)
}

func (a *AuthHandler) ValidateJWT(c echo.Context) error {
	token := c.Param("token")

	err := a.AuthService.ValidateJWT(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &erToken{ErrorToken: err.Error()})
	}

	return c.JSON(http.StatusCreated, "valid token")
}

type erToken struct {
	ErrorToken string `json:"error"`
}