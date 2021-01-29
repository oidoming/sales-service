package auth

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)


func (a *AuthHandler) MiddlewareValidateAccessJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken, _ := extractToken(c)

		err := a.AuthService.ValidateJWT(accessToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, &erToken{ErrorToken: err.Error()})
		}

		log.Println("valid token")

		return next(c)
	}
}

func (a *AuthHandler) MiddlewareValidateRefreshJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, refreshToken := extractToken(c)

		userId, err := a.AuthService.ValidateRefreshJWT(refreshToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, &erToken{ErrorToken: err.Error()})
		}

		log.Println("valid refresh token")

		c.Set("user_id", userId)

		return next(c)
	}
}

func extractToken(c echo.Context) (accessJWT, refreshJWT string){
	accessJWT = c.Request().Header.Get("x-access-token")
	refreshJWT = c.Request().Header.Get("x-refresh-token")

	return accessJWT, refreshJWT
}
