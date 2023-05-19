package user

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestUserService_GetUserById functionality
func TestUserService_GetUserById(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchedUser, err := service.GetUserById(users[0].Id)
	assertUsersEquality(t, fetchedUser, &users[0])

	randId := rand.Int()
	_, err = service.GetUserById(randId)
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")
}

// TestUserService_GetUserByUUID functionality
func TestUserService_GetUserByUUID(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchedUser, err := service.GetUserByUUID(users[0].UUID)
	assertUsersEquality(t, fetchedUser, &users[0])

	randUUID := "Test2UUid"
	_, err = service.GetUserByUUID(randUUID)
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")
}

// TestUserService_GetUserByPhone functionality
func TestUserService_GetUserByPhone(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	fetchedUser, err := service.GetUserByPhone(users[0].PhoneNumber)
	assertUsersEquality(t, fetchedUser, &users[0])

	randPhone := "09191234567"
	_, err = service.GetUserByPhone(randPhone)
	assert.Error(t, err, "Fetching wrong user from db failed ! it should throw an error")
}

// TestUserService_CreateUser functionality
func TestUserService_CreateUser(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	mockedUser := mockUser()

	createdUser, err := service.CreateUser(mockedUser.Name, mockedUser.PhoneNumber, mockedUser.Password)
	defer destructCreatedObjects(db, []User{*createdUser})

	assert.NoError(t, err, "User service user creation failed")
	assert.NotEqual(t, createdUser.Password, mockedUser.Password, "User service user creation failed")
	assert.NotZero(t, createdUser.Id, "User service user creation failed")

	_, err = service.CreateUser(mockedUser.Name, mockedUser.PhoneNumber, mockedUser.Password)
	assert.Error(t, err, "User service user creation failed")
	assert.ErrorIs(t, err, UserAlreadyExists, "User service user creation failed")
}

// TestUserService_UpdateUser functionality
func TestUserService_UpdateUser(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Setup database connection failed")

	service := createService(db)
	users := mockAndInsertUser(db, 1)
	defer destructCreatedObjects(db, users)

	newName := "newlyyyname"
	pass := "Password"

	updatedUser, err := service.UpdateUser(users[0].Id, &newName, &pass)
	assert.NoError(t, err, "User service update user failed")
	assert.Equal(t, *updatedUser.Name, newName, "User service update user failed")
	assert.NotEqual(t, pass, updatedUser.Password, "User service update user failed")

	randId := rand.Int()
	_, err = service.UpdateUser(randId, &newName, &pass)
	assert.Error(t, err, "User service update user failed")
	assert.ErrorIs(t, err, UserDoesntExists, "User service update user failed")
}

// createService and return it for testing purpose
func createService(db *gorm.DB) UserServiceInterface {
	return NewService(NewRepository(db))
}
