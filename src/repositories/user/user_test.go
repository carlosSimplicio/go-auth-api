package userRepository

import (
	"testing"

	interfaces "github.com/carlosSimplicio/go-auth-api/src/registry"
)

func TestCreateUser(t *testing.T) {
	testUser := interfaces.User{
		Id:       1,
		Name:     "i am mocked",
		Email:    "mock@mock.com",
		Password: "mockingIsFun",
	}
}
