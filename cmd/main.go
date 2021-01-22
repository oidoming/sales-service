package main

import (
	"fmt"

	"github.com/Oscar-inc117/sales-service/api/handlers"
	"github.com/Oscar-inc117/sales-service/internal/config"
	"github.com/Oscar-inc117/sales-service/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appConfig := config.Load()
	port := appConfig.Server.Port

	client := &handlers.Client{
		DB: database.GetConnection(&appConfig.Database),
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.POST("/api/clients", client.CreateClient)
	e.GET("/api/clients", client.GetClients)
	e.GET("/api/clients/:id", client.GetClient)
	e.PUT("/api/clients/:id", client.UpdateClient)
	e.DELETE("/api/clients/:id", client.DeleteClient)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
