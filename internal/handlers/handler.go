package handlers

import (
	"github.com/Oscar-inc117/sales-service/internal/handlers/auth"
	"github.com/Oscar-inc117/sales-service/internal/handlers/client"
	"github.com/Oscar-inc117/sales-service/internal/handlers/product"
	"github.com/Oscar-inc117/sales-service/internal/handlers/sales"
	"github.com/Oscar-inc117/sales-service/internal/services/authsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/clientsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/productsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/salessrv"
	"github.com/Oscar-inc117/sales-service/internal/services/usersrv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(
	clientSrv clientsrv.Service,
	userSrv usersrv.Service,
	authSrv authsrv.Service,
	salesSrv salessrv.Service,
	productSrv productsrv.Service) *echo.Echo {

	ah := auth.NewAuthHandler(authSrv, userSrv)
	ch := client.NewClientHandler(clientSrv)
	ph := product.NewProductHandler(productSrv)
	sh := sales.NewSalesHandler(salesSrv)

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

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

	g.POST("/products", ph.AddProduct)
	g.GET("/products", ph.GetProducts)
	g.GET("/products/:id", ph.GetProduct)
	g.PUT("/products/:id", ph.UpdateProduct)
	g.DELETE("/products/:id", ph.DeleteProduct)

	g.POST("/sales", sh.NewSale)
	g.GET("/sales", sh.GetSales)
	g.GET("/sales/:id", sh.GetSale)

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
