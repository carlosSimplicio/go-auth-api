package services

import (
	"fmt"

	userRepository "github.com/carlosSimplicio/go-auth-api/src/repositories/user"
	"github.com/carlosSimplicio/go-auth-api/src/utils"
)

type SignUpData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(body []byte) error {
	signUpData := &SignUpData{}
	err := utils.ParseJson(body, signUpData)

	if err != nil {
		return fmt.Errorf("Failed to parse body: [%w]", err)
	}

	hashedPassword, err := utils.HashPassword([]byte(signUpData.Password))

	if err != nil {
		return fmt.Errorf("Error hashing password: [%w]", err)
	}

	user := userRepository.User{
		Name:     signUpData.Name,
		Email:    signUpData.Email,
		Password: string(hashedPassword),
	}

	_, err = userRepository.CreateUser(&user)

	if err != nil {
		return fmt.Errorf("Failed to create user: [%w]", err)
	}

	return nil
}

func Login(body []byte) (string, error) {
	loginData := &LoginData{}
	err := utils.ParseJson(body, loginData)

	if err != nil {
		return "", fmt.Errorf("Failed to parse body: [%w]", err)
	}

	user, err := userRepository.GetUserByEmail(loginData.Email)

	if err != nil {
		return "", fmt.Errorf("User not found: [%w]", err)
	}

	err = utils.ComparePassword([]byte(user.Password), []byte(loginData.Password))

	if err != nil {
		return "", fmt.Errorf("Invalid password", err)
	}

	jwtToken, err := utils.CreateJwt(user)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
