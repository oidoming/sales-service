package handlers

import (
	"github.com/Oscar-inc117/sales-service/internal/handlers/client"
	"github.com/Oscar-inc117/sales-service/internal/handlers/user"
	"github.com/Oscar-inc117/sales-service/internal/services/clientsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/usersrv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(clientSrv clientsrv.Service, userSrv usersrv.Service) *echo.Echo {
	e := echo.New()

	e.POST("/api/clients", client.CreateClient(clientSrv))
	e.GET("/api/clients", client.GetClients(clientSrv))
	e.GET("/api/clients/:id", client.GetClient(clientSrv))
	e.PUT("/api/clients/:id", client.UpdateClient(clientSrv))
	e.DELETE("/api/clients/:id", client.DeleteClient(clientSrv))

	e.POST("/api/user", user.CreateUser(userSrv))
	e.GET("/api/user", user.GetUser(userSrv))
	e.PUT("/api/user/:id", user.UpdateUser(userSrv))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	return e
}
