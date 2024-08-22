package authentication

import (
	"errors"
	"fmt"

	userRepository "github.com/carlosSimplicio/go-auth-api/src/repositories/user"
	"github.com/carlosSimplicio/go-auth-api/src/utils"
)

type SignUpData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(body []byte) error {
	signUpData := &SignUpData{}
	err := utils.ParseJson(body, signUpData)

	if err != nil {
		return fmt.Errorf("failed to parse body: [%w]", err)
	}

	if signUpData.Email == "" || signUpData.Name == "" || signUpData.Password == "" {
		return errors.New("invalid missing required params: email, name or password")
	}

	user, err := userRepository.GetUserByEmail(signUpData.Email)

	if err != nil && !errors.Is(userRepository.ErrUserNotFound, err) {
		return err
	}

	if user != nil {
		return errors.New("email already in use")
	}

	hashedPassword, err := utils.HashPassword([]byte(signUpData.Password))

	if err != nil {
		return fmt.Errorf("error hashing password: [%w]", err)
	}

	user = &userRepository.User{
		Name:     signUpData.Name,
		Email:    signUpData.Email,
		Password: string(hashedPassword),
	}

	_, err = userRepository.CreateUser(user)

	if err != nil {
		return fmt.Errorf("failed to create user: [%w]", err)
	}

	return nil
}
