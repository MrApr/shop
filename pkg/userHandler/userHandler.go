package userHandler

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"shop/internal/domains/user"
	"shop/pkg/database"
	"shop/pkg/tokenizer"
	"sync"
)

var (
	UserNotFoundErr    error = fmt.Errorf("%s", "Requested User not found for fetching ID")
	AuthSomethingWrong error = fmt.Errorf("%s", "Something went wrong, please contact administrator")
)
var repo user.UserRepositoryInterface
var once sync.Once

// init initializes user handler
func initFunc() {
	db := connectToDb()
	repo = user.NewRepository(db)
}

// connectToDb makes database connection
func connectToDb() *gorm.DB {
	conn, err := database.Conn()
	if err != nil {
		panic(err)
	}
	return conn
}

// ExtractUserIdFromToken which is apssed
func ExtractUserIdFromToken(ctx context.Context, token string) (int, error) {
	once.Do(initFunc)
	tokenGenerator := tokenizer.CreateTokenizer(ctx)
	tkInfo, err := tokenGenerator.TokenInfo(token)
	if err != nil {
		return 0, AuthSomethingWrong
	}

	fetchedUserId, err := getUserId(tkInfo.UUID)
	if err != nil {
		return 0, err
	}

	return fetchedUserId, nil
}

// getUserId from user uuid and return it
func getUserId(uuid string) (int, error) {
	fetchedUser, err := repo.GetUserByUUID(uuid)
	if err != nil {
		return 0, UserNotFoundErr
	}
	return fetchedUser.Id, nil
}
