package userRepository

import (
	"errors"
	"fmt"

	"github.com/carlosSimplicio/go-auth-api/src/controllers"
	"github.com/carlosSimplicio/go-auth-api/src/infra/mysql"
)

func CreateUser(user *controllers.User) (int, error) {
	query := "INSERT INTO user (name, email, password) VALUES (?,?,?);"
	params := []any{user.Name, user.Email, user.Password}

	result, err := mysql.Exec(query, params...)

	if err != nil {
		return 0, fmt.Errorf("Failed to insert user: [%w]", err)
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("Failed to get inserted user ID: [%w]", err)
	}

	return int(userId), nil
}

func GetUserById(id int) (*controllers.User, error) {
	query := "SELECT Id, Name, Email, Password FROM user WHERE id = ?;"
	result, err := mysql.Select[controllers.User](query, id)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch User: [%w]", err)
	}

	if len(result) == 0 {
		return nil, errors.New("User not found")
	}

	user := result[0]

	return &user, nil
}

func GetUserByEmail(email string) (*controllers.User, error) {
	query := "SELECT Id, Name, Email, Password FROM user WHERE email = ?;"
	result, err := mysql.Select[controllers.User](query, email)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch User: [%w]", err)
	}

	if len(result) == 0 {
		return nil, errors.New("User not found")
	}

	user := result[0]

	return &user, nil
}
