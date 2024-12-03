package userRepository

import (
	"fmt"
	"testing"

	interfaces "github.com/carlosSimplicio/go-auth-api/src/registry"
)

var mockClient interfaces.DbClient

func TestCreateUser(t *testing.T) {
	testUser := interfaces.User{
		Name:     "i am mocked",
		Email:    "mock@mock.com",
		Password: "mockingIsFun",
	}

	repository := UserRepository{
		Client: mockClient,
	}

	insertedId, err := repository.CreateUser(&testUser)

	fmt.Sprintf("%d, %s", insertedId, err)

}
