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

type IDbClient interface {
	Connect()
	Select(query string, params ...any) (*sql.Rows, error)
	Exec(query string, params ...any) (operationResult sql.Result, err error)
	Close()
}

type IUserRepository interface {
	CreateUser(user *User) (int, error)
	GetUserById(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type ILoginService interface {
	Login(body []byte) (string, error)
}

type ISignUpService interface {
	SignUp(body []byte) error
}
