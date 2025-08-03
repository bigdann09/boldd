package jwt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/boldd/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	GenerateAccessToken(id int, email string) string
	GenerateRefreshToken(id int) string
	ValidateToken(token string) (*Claims, error)
}

type TokenService struct {
	key            string
	access_expiry  int
	refresh_expiry int
}

type Claims struct {
	Id    int
	Email string
	jwt.RegisteredClaims
}

func NewTokenService(cfg *config.JWTConfig) *TokenService {
	return &TokenService{
		key:            cfg.Key,
		access_expiry:  cfg.AccessExpiry,
		refresh_expiry: cfg.RefreshExpiry,
	}
}

func (srv TokenService) GenerateAccessToken(id int, email string) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(id),
			Issuer:    strconv.Itoa(id),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(id),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(srv.access_expiry))),
		},
	})

	token, _ := claims.SignedString([]byte(srv.key))
	return token
}

func (srv TokenService) GenerateRefreshToken(id int) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        strconv.Itoa(id),
		Issuer:    strconv.Itoa(id),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   strconv.Itoa(id),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(srv.refresh_expiry))),
	})

	token, _ := claims.SignedString([]byte(srv.key))
	return token
}

func (srv TokenService) ValidateToken(token_string string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(token_string, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(srv.key), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, fmt.Errorf("token expired")
	}
	return claims, nil
}
