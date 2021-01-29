package main

import (
	"fmt"
	"github.com/Oscar-inc117/sales-service/internal/config"
	"github.com/Oscar-inc117/sales-service/internal/database"
	"github.com/Oscar-inc117/sales-service/internal/handlers"
	"github.com/Oscar-inc117/sales-service/internal/repository/postgres"
	"github.com/Oscar-inc117/sales-service/internal/services/authsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/clientsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/usersrv"
)

func main() {
	appConfig := config.Load()
	port := appConfig.Server.Port

	clientRepo := postgres.ClientStorage{
		DB: database.GetConnection(&appConfig.Database),
	}
	clientService := clientsrv.NewService(&clientRepo)

	userRepo := postgres.UserRepo{
		DB: database.GetConnection(&appConfig.Database),
	}
	userService := usersrv.NewService(&userRepo)
	authService := authsrv.NewService(&userRepo)

	e := handlers.Handler(clientService, userService, authService)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

