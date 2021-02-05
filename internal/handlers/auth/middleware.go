package auth

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)


func (a *AuthHandler) MiddlewareValidateAccessJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			log.Println(err)
		}

		err = a.AuthService.ValidateJWT(cookie.Value)
		if err != nil {
			cookie.Value = ""
			cookie.MaxAge = -1
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
	cookie, err := c.Cookie("access_token")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)

	refreshJWT = c.Request().Header.Get("x-refresh-token")

	return cookie.Value, refreshJWT
}
