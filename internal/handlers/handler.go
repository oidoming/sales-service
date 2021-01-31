package handlers

import (
	"github.com/Oscar-inc117/sales-service/internal/handlers/auth"
	"github.com/Oscar-inc117/sales-service/internal/handlers/client"
	"github.com/Oscar-inc117/sales-service/internal/services/authsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/clientsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/usersrv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(clientSrv clientsrv.Service, userSrv usersrv.Service, authSrv authsrv.Service) *echo.Echo {
	ch := client.NewClientHandler(clientSrv)
	ah := auth.NewAuthHandler(authSrv, userSrv)

	e := echo.New()

	e.POST("/api/auth", ah.CreateUser)
	e.GET("/api/auth", ah.GetUser)
	e.PUT("/api/auth/:id", ah.UpdateUser)

	e.POST("/api/login", ah.Login)
	e.GET("/api/:token", ah.ValidateJWT)

	r := e.Group("api/refresh")
	r.Use(ah.MiddlewareValidateRefreshJWT)
	r.POST("", ah.RefreshToken)

	//e.Use(ah.MiddlewareValidateAccessJWT)
	g := e.Group("/api/admin")
	g.Use(ah.MiddlewareValidateAccessJWT)
	g.POST("/clients", ch.CreateClient)
	g.GET("/clients", ch.GetClients)
	g.GET("/clients/:id", ch.GetClient)
	g.PUT("/clients/:id", ch.UpdateClient)
	g.DELETE("/clients/:id", ch.DeleteClient)


	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderSetCookie,
			"*"},
	}))

	return e
}
