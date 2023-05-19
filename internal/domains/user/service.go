package user

import (
	"shop/pkg/advancedError"
	"shop/pkg/hasher"
	"shop/pkg/uuid"
)

// UserService which implements service interface functionalities and satisfies it
type UserService struct {
	repo UserRepositoryInterface
}

// NewService instantiates and returns new user service
func NewService(repo UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

// GetUserById and return it
func (u *UserService) GetUserById(id int) (*User, error) {
	return u.repo.GetUserById(id)
}

// GetUserByUUID and return it
func (u *UserService) GetUserByUUID(uuid string) (*User, error) {
	return u.repo.GetUserByUUID(uuid)
}

// GetUserByPhone and return it
func (u *UserService) GetUserByPhone(phone string) (*User, error) {
	return u.repo.GetUserByPhone(phone)
}

// CreateUser that doesn't exist in repository
func (u *UserService) CreateUser(name *string, phone, password string) (*User, error) {
	isExists := u.repo.UserExists(phone)
	if isExists {
		return nil, UserAlreadyExists
	}

	hashedPass, err := hasher.Make(password)
	if err != nil {
		return nil, advancedError.New(err, "user hashing password failed")
	}

	generatedUUID, err := uuid.GenerateUUId()
	if err != nil {
		return nil, advancedError.New(err, "user generate uuid failed")
	}

	user := &User{
		Id:          0,
		UUID:        generatedUUID,
		PhoneNumber: phone,
		Name:        name,
		Password:    hashedPass,
	}

	return u.repo.CreateUser(user)
}

// UpdateUser that already exists in database
func (u *UserService) UpdateUser(userId int, name, password *string) (*User, error) {
	user, err := u.repo.GetUserById(userId)
	if err != nil {
		return nil, UserDoesntExists
	}

	if password != nil {
		password, err = u.makePasswordUseAble(password)
	}

	if err != nil {
		return nil, advancedError.New(err, "Cannot hash password")
	}

	return u.repo.UpdateUser(user, name, password)
}

// makePasswordUseAble for in update operation
func (u *UserService) makePasswordUseAble(password *string) (*string, error) {
	hashedPass, err := hasher.Make(*password)
	if err != nil {
		return nil, err
	}

	return &hashedPass, nil
}
