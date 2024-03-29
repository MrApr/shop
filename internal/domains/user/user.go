package user

import "context"

// UserRepositoryInterface defines set of methods which every repository who wants to play role as user repo should obey.
type UserRepositoryInterface interface {
	GetUserById(id int) (*User, error)
	GetUserByUUID(uuid string) (*User, error)
	GetUserByPhone(phone string) (*User, error)
	UserExists(phone string) bool
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User, name, password *string) (*User, error)
}

// UserServiceInterface defines set of methods which every service who wants to play role as user service should obey.
type UserServiceInterface interface {
	GetUserById(id int) (*User, error)
	GetUserByUUID(uuid string) (*User, error)
	GetUserByPhone(phone string) (*User, error)
	CreateUser(name *string, phone, password string) (*User, error)
	UpdateUser(userId int, name, password *string) (*User, error)
}

// UserUseCaseInterface defines set of methods which every use case who wants to play role as user use case should obey.
type UserUseCaseInterface interface {
	Register(ctx context.Context, request *UserRegisterRequest) (*AuthResponse, error)
	Login(ctx context.Context, request *UserLoginRequest) (*AuthResponse, error)
	UpdateUserPass(ctx context.Context, token string, request *UpdateUserPasswordRequest) (*User, error)
	UpdateUserName(ctx context.Context, token string, request *UpdateUserRequest) (*User, error)
}
