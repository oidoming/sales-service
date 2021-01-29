package auth

import (
	"github.com/Oscar-inc117/sales-service/internal/services/authsrv"
	"github.com/Oscar-inc117/sales-service/internal/services/usersrv"
)
type AuthHandler struct {
	AuthService authsrv.Service
	UserService usersrv.Service
}

func NewAuthHandler(auth authsrv.Service, user usersrv.Service) *AuthHandler {
	return &AuthHandler{
		AuthService: auth,
		UserService: user,
	}
}