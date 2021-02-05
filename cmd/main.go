package main

import (
	"fmt"
	"github.com/Oscar-inc117/sales-service/internal/config"
	"github.com/Oscar-inc117/sales-service/internal/database"
	"github.com/Oscar-inc117/sales-service/internal/handlers"
	"github.com/Oscar-inc117/sales-service/internal/repository/postgres"
	"github.com/Oscar-inc117/sales-service/internal/services/authsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/clientsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/productsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/salessrv"
	"github.com/Oscar-inc117/sales-service/internal/services/usersrv"
)

func main() {
	appConfig := config.Load()
	port := appConfig.Server.Port

	repo := postgres.NewRepository(database.GetConnection(&appConfig.Database))

	userService := usersrv.NewService(&repo)
	authService := authsrv.NewService(&repo)
	salesService := salessrv.NewService(&repo)
	clientService := clientsrv.NewService(&repo)
	productService := productsrv.NewService(&repo)

	e := handlers.Handler(clientService, userService, authService, salesService, productService)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

