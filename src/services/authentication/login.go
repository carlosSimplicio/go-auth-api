package authentication

import (
	"fmt"

	userRepository "github.com/carlosSimplicio/go-auth-api/src/repositories/user"
	"github.com/carlosSimplicio/go-auth-api/src/utils"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(body []byte) (string, error) {
	loginData := &LoginData{}
	err := utils.ParseJson(body, loginData)

	if err != nil {
		return "", fmt.Errorf("failed to parse body: [%w]", err)
	}

	user, err := userRepository.GetUserByEmail(loginData.Email)

	if err != nil {
		return "", fmt.Errorf("user not found: [%w]", err)
	}

	err = utils.ComparePassword([]byte(user.Password), []byte(loginData.Password))

	if err != nil {
		return "", fmt.Errorf("invalid password: [%w]", err)
	}

	jwtToken, err := utils.CreateJwt(user)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}