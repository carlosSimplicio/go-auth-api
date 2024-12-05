package userRepository

import (
	"database/sql"
	"testing"

	interfaces "github.com/carlosSimplicio/go-auth-api/src/registry"
	mock_interfaces "github.com/carlosSimplicio/go-auth-api/src/testing"
	"go.uber.org/mock/gomock"
)

type mockExecReturn struct{}

var execLastInsertId int64
var execRowsAffected int64
var execError error

func (mockExecReturn) LastInsertId() (int64, error) {
	return execLastInsertId, execError
}

func (mockExecReturn) RowsAffected() (int64, error) {
	return execRowsAffected, execError
}

func TestCreateUserSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mock_interfaces.NewMockIDbClient(ctrl)
	testUser := &interfaces.User{
		Name:     "i am mocked",
		Email:    "mock@mock.com",
		Password: "mockingIsFun",
	}
	execLastInsertId = 25

	mockClient.EXPECT().Exec(
		gomock.Eq("INSERT INTO user (name, email, password) VALUES (?,?,?);"),
		gomock.Eq([]any{testUser.Name, testUser.Email, testUser.Password}),
	).Times(1).Return(sql.Result(mockExecReturn{}), nil)

	repository := &UserRepository{
		Client: mockClient,
	}

	insertedId, err := repository.CreateUser(testUser)
	if err != nil {
		t.Fatalf("Error creating user: %s\n", err)
	}

	if insertedId != int(execLastInsertId) {
		t.Fatalf("Exepected insertedId: %d, Got: %d!\n", execLastInsertId, insertedId)
	}

}
