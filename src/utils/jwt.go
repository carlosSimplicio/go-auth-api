package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/carlosSimplicio/go-auth-api/src/controllers"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJwt(user *controllers.User) (string, error) {
	type Claims struct {
		userId    int
		userName  string
		userEmail string
		jwt.RegisteredClaims
	}
	registeredClaims := jwt.RegisteredClaims{
		Issuer:    "go-auth-api",
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 7)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		user.Id, user.Name, user.Email, registeredClaims,
	})

	key := "meu-segredo"
	s, err := token.SignedString([]byte(key))

	if err != nil {
		return "", fmt.Errorf("Failed to create jwt: [%w]", err)
	}

	return s, nil
}

func VerifyJwt(tokenStr string) {
	parser := jwt.Parser{}
	token, err := parser.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte("meu-segredo"), nil
	})

	if err != nil {
		log.Fatalf("Failed to verify token", err)
	}

	fmt.Printf("%v", token.Valid)
}
