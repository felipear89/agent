package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/felipear89/agent/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type Service struct {
	cfg *config.AuthConfig
}

func newService(cfg *config.AuthConfig) *Service {
	return &Service{
		cfg: cfg,
	}
}

type Token struct {
	*jwt.Token

	ExpiresAt time.Time
	SignedJWT string
}

func (a *Service) GenerateToken(userID int, email string) (*Token, error) {
	expiresAt := time.Now().Add(a.cfg.TokenExpiryDuration())
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        a.cfg.TokenID,
			Issuer:    a.cfg.Issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, err := token.SignedString(a.cfg.JWTPrivateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %w", err)
	}
	return &Token{
		Token:     token,
		ExpiresAt: expiresAt,
		SignedJWT: signedString,
	}, nil
}

func (a *Service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.cfg.JWTPublicKeyPEM, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
