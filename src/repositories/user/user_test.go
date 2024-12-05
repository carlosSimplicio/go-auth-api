package userRepository

import (
	"fmt"
	"testing"

	interfaces "github.com/carlosSimplicio/go-auth-api/src/registry"
	mock_interfaces "github.com/carlosSimplicio/go-auth-api/src/testing"
	"go.uber.org/mock/gomock"
)

type mockExecReturn struct {
	returnId int64
}

func (m *mockExecReturn) LastInsertId() (int64, error) {
	return m.returnId, nil
}

func TestCreateUserSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mock_interfaces.NewMockIDbClient(ctrl)
	testUser := &interfaces.User{
		Name:     "i am mocked",
		Email:    "mock@mock.com",
		Password: "mockingIsFun",
	}
	mockReturnId := int64(25)

	mockClient.EXPECT().Exec(
		gomock.Eq("INSERT INTO user (name, email, password) VALUES (?, ?, ?);"),
		gomock.Eq([]any{testUser.Name, testUser.Email, testUser.Password}),
	).Times(1).Return(mockExecReturn{returnId: mockReturnId}, nil)

	repository := &UserRepository{
		Client: mockClient,
	}

	insertedId, err := repository.CreateUser(testUser)

	fmt.Sprintf("%d, %s", insertedId, err)

}
