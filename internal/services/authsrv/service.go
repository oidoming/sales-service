package authsrv

import (
	"errors"
	"github.com/Oscar-inc117/sales-service/internal/domain"
	"github.com/Oscar-inc117/sales-service/internal/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"log"
	"time"
)

const secretKeyJWT string = "alajaja4"
const secretKeyRefresh string = "wala949h"

type Service interface {
	Auth(email, password string) (domain.User, bool)
	GenerateJWT(user domain.User) string
	GenerateRefreshJWT(user domain.User) string
	ValidateJWT(tokenString string) error
	ValidateRefreshJWT(tokenString string) (userId string, err error)
}

type Repository interface {
	SelectUserByEmail(email string) (domain.User, error)
}

type authService struct {
	r Repository
}

func NewService(r Repository) Service {
	return &authService{r}
}

func (a *authService) Auth(email, password string) (domain.User, bool) {
	user, err := a.r.SelectUserByEmail(email)
	if err != nil {
		log.Println(err)
		return user, false
	}

	samePassword := services.ComparePassword(password, user.Password)

	if email != user.Email || samePassword == false {
		return user, false
	}

	return user, true
}

func (a *authService) GenerateJWT(user domain.User) string {
	claims := jwt.MapClaims{
		"access_id": uuid.New(),
		"user_id": user.ID.String(),
		"iss": "sales-service",
		"exp": time.Now().Add(time.Minute * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, _ := token.SignedString([]byte(secretKeyJWT))

	return tokenString
}

func (a *authService) GenerateRefreshJWT(user domain.User) string {
	claims := jwt.MapClaims{
		"refresh_id": uuid.New(),
		"user_id": user.ID.String(),
		"iss": "sales-service",
		"exp": time.Now().Add(time.Hour * 42 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, _ := token.SignedString([]byte(secretKeyRefresh))

	return tokenString
}

func (a *authService) ValidateJWT(tokenString string) error {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKeyJWT), nil
	})
	if err != nil {
		log.Println("validate", err)
		return err
	}

	_, ok := token.Claims.(jwt.Claims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func (a *authService) ValidateRefreshJWT(tokenString string) (userId string, err error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKeyRefresh), nil
	})
	if err != nil {
		log.Println("validate", err)
		return "", err
	}

	_, ok := token.Claims.(jwt.Claims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims["user_id"].(string), nil
}
/*
func (a *authService) ParseToken(tokenString string) (userId string) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKeyJWT), nil
	})
	if err != nil {
		log.Println("parse ", err)
		return 
	}

	return claims["user_id"].(string)
}
*/
