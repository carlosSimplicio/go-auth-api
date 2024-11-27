package interfaces

import (
	"database/sql"
	"net/http"
)

type Controller interface {
	SetupRoutes(handler *http.ServeMux)
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type DbClient interface {
	Connect()
	Select(query string, params ...any) (*sql.Rows, error)
	Exec(query string, params ...any) (operationResult sql.Result, err error)
}

type UserRepository interface {
	CreateUser(user *User) (int, error)
	GetUserById(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type LoginService interface {
	Login(body []byte) (string, error)
}

type SignUpService interface {
	SignUp(body []byte) error
}
