package user

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"shop/pkg/tokenizer"
	"testing"
)

// TestUserUseCase_Login functionality
func TestUserUseCase_Login(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	ctx := context.WithValue(context.Background(), ContextValueIpKey, "192.168.1.1")
	usecase := createUseCase(db)
	sv := createService(db)

	mockedLoginReq := mockLoginRequest()
	createdUser, err := sv.CreateUser(nil, mockedLoginReq.PhoneNumber, mockedLoginReq.Password)
	assert.NoError(t, err, "Use creation in test user use-case failed")

	result, err := usecase.Login(ctx, mockedLoginReq)
	fmt.Println(result.Tokens)
	assert.NoError(t, err, "login failed in user use-case")
	assert.NotNil(t, result.Tokens, "login failed in user use-case")
	assert.NotNil(t, result.User, "login failed in user use-case")
	assertUsersEquality(t, result.User, createdUser)
}

// TestUserUseCase_Register functionality
func TestUserUseCase_Register(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	useCase := createUseCase(db)
	registerReq := mockRegisterRequest()
	ctx := context.WithValue(context.Background(), ContextValueIpKey, "192.168.1.1")

	result, err := useCase.Register(ctx, registerReq)
	assert.NoError(t, err, "register failed in user use-case")
	assert.NotNil(t, result.Tokens, "register failed in user use-case")
	assert.NotNil(t, result.User, "register failed in user use-case")
	assert.Equal(t, *result.User.Name, *registerReq.Name, "register failed in user use-case")
}

// TestUserUseCase_UpdateUserName functionality
func TestUserUseCase_UpdateUserName(t *testing.T) {

}

// TestUserUseCase_UpdateUserPass functionality
func TestUserUseCase_UpdateUserPass(t *testing.T) {

}

// createUseCase and return it for testing purpose
func createUseCase(db *gorm.DB) UserUseCaseInterface {
	return NewUserUseCase(NewService(NewRepository(db)))
}

// mockLoginRequest and return it for login operation
func mockLoginRequest() *UserLoginRequest {
	return &UserLoginRequest{
		PhoneNumber: "09191234567",
		Password:    "1234567879",
	}
}

// mockRegisterRequest and return it for register operation
func mockRegisterRequest() *UserRegisterRequest {
	name := "KingApr"
	return &UserRegisterRequest{
		PhoneNumber:     "09191234567",
		Password:        "1234567879",
		PasswordConfirm: "1234567879",
		Name:            &name,
	}
}

// mockUserJwtTk and return it
func mockUserJwtTk(ctx context.Context, uuid, ip string) (string, error) {
	tokenGenerator := tokenizer.CreateTokenizer(ctx)
	jwtTk, err := tokenGenerator.New(uuid, ip)
	if err != nil {
		return "", err
	}

	return jwtTk["token"], nil
}
