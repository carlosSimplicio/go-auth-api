package userRepository

import (
	"errors"
	"fmt"

	interfaces "github.com/carlosSimplicio/go-auth-api/src/registry"
	"github.com/carlosSimplicio/go-auth-api/src/utils"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	Client interfaces.DbClient
}

func (u *UserRepository) CreateUser(user *interfaces.User) (int, error) {
	query := "INSERT INTO user (name, email, password) VALUES (?,?,?);"
	params := []any{user.Name, user.Email, user.Password}

	result, err := u.Client.Exec(query, params...)

	if err != nil {
		return 0, fmt.Errorf("failed to insert user: [%w]", err)
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("failed to get inserted user ID: [%w]", err)
	}

	return int(userId), nil
}

func (u *UserRepository) GetUserById(id int) (*interfaces.User, error) {
	query := "SELECT Id, Name, Email, Password FROM user WHERE id = ?;"
	rows, err := u.Client.Select(query, id)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch User: [%w]", err)
	}

	result, err := utils.GetRowsValues[interfaces.User](rows)

	if err != nil {
		return nil, fmt.Errorf("failed to get values from rows for User: [%w]", err)
	}

	if len(result) == 0 {
		return nil, ErrUserNotFound
	}

	user := result[0]

	return &user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (*interfaces.User, error) {
	query := "SELECT Id, Name, Email, Password FROM user WHERE email = ?;"
	rows, err := u.Client.Select(query, email)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch User: [%w]", err)
	}

	result, err := utils.GetRowsValues[interfaces.User](rows)

	if err != nil {
		return nil, fmt.Errorf("failed to get values from rows for User: [%w]", err)
	}

	if len(result) == 0 {
		return nil, ErrUserNotFound
	}

	user := result[0]

	return &user, nil
}
