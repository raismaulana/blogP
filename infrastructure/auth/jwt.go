package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
)

type JWTToken struct {
	env *envconfig.EnvConfig
}

type customClaims struct {
	authorized bool
	email      string
	role       string
	jwt.StandardClaims
}

func NewJWTToken(env *envconfig.EnvConfig) (*JWTToken, error) {
	if strings.TrimSpace(env.SecretKey) == "" {
		return nil, fmt.Errorf("JWT secret key must not empty")
	}

	return &JWTToken{
		env: env,
	}, nil

}

func (r *JWTToken) GenerateToken(id_user string, email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &customClaims{
		authorized: true,
		email:      email,
		role:       role,
		StandardClaims: jwt.StandardClaims{
			Audience:  r.env.AppBaseURL,
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
			Id:        uuid.NewString(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    r.env.AppBaseURL,
			NotBefore: time.Now().Unix(),
			Subject:   id_user,
		},
	})

	tokenString, err := token.SignedString([]byte(r.env.SecretKey))
	if err != nil {
		return "", err

	}
	return tokenString, nil
}

func (r *JWTToken) VerifyToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])

		}
		return []byte(r.env.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
