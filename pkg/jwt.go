package pkg

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ErrmissingKey = errors.New("jwt key not found")

type JWTClaims struct {
	Id   int
	Role string
	jwt.RegisteredClaims
}

func NewJWTClaims(id int, role string) *JWTClaims {
	ID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &JWTClaims{
		Id:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			ID:        ID.String(),
		},
	}
}

func (j *JWTClaims) GenToken() (string, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return "", ErrmissingKey
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))

}

func (j *JWTClaims) VerifyToken(token string) error {
	jwtToken, err := jwt.ParseWithClaims(token, j, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if !jwtToken.Valid {
		return jwt.ErrTokenExpired
	}
	issuer, err := jwtToken.Claims.GetIssuer()
	if err != nil {
		return err
	}
	if issuer != os.Getenv("JWT_ISSUER") {
		return jwt.ErrTokenInvalidIssuer
	}
	return nil
}
