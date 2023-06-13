package paymentHandler

import (
	"errors"
	"gorm.io/gorm"
	"shop/internal/domains/address"
	"shop/internal/domains/basket"
	"shop/internal/domains/discount"
	"shop/internal/domains/gateways"
	"shop/internal/domains/postType"
	"shop/pkg/database"
	"strings"
	"sync"
)

var (
	BasketIsInvalid   error = errors.New("selected basket is invalid")
	AddressIsInvalid  error = errors.New("selected address is invalid")
	DiscountIsInvalid error = errors.New("selected discount is invalid")
	PostTypeIsInvalid error = errors.New("selected post type is invalid")
)

var basketRepo basket.BasketRepositoryInterface
var addressRepo address.AddressRepositoryInterface
var discountRepo discount.DiscountRepositoryInterface
var postTypeRepo postType.PostTypeRepositoryInterface
var gatewayRepo gateways.GatewayRepositoryInterface
var once sync.Once

// init initializes user handler
func initFunc() {
	db := connectToDb()
	basketRepo = basket.NewBasketRepository(db)
	addressRepo = address.NewAddressRepository(db)
	discountRepo = discount.NewRepository(db)
	postTypeRepo = postType.NewRepository(db)
	gatewayRepo = gateways.NewGatewayRepository(db)
}

// connectToDb makes database connection
func connectToDb() *gorm.DB {
	conn, err := database.Conn()
	if err != nil {
		panic(err)
	}
	return conn
}

// BasketIsValid checks whether given basket exists or not
func BasketIsValid(id int) error {
	once.Do(initFunc)
	exists := basketRepo.BasketExists(id)
	if !exists {
		return BasketIsInvalid
	}
	return nil
}

// AddressIsValid checks whether given address is valid or not
func AddressIsValid(id int) error {
	once.Do(initFunc)
	_, err := addressRepo.GetAddressById(id)
	if err != nil {
		return AddressIsInvalid
	}
	return nil
}

// IsDiscountValid checks whether given discount is valid or not
func IsDiscountValid(id int) error {
	once.Do(initFunc)
	discountCode := discountRepo.GetDiscountById(id)
	if discountCode.Id == 0 {
		return DiscountIsInvalid
	}
	return nil
}

// IsPostTypeValid checks whether post type is valid or not
func IsPostTypeValid(id int) error {
	exists := postTypeRepo.PostTypeExists(id)
	if !exists {
		return PostTypeIsInvalid
	}
	return nil
}

// GetGatewayToken and return it
func GetGatewayToken(requestedGwTitle string) string {
	gateways := gatewayRepo.GetAllGateways(1, true)
	for _, gateway := range gateways {
		if strings.ToLower(gateway.Name) == strings.ToLower(requestedGwTitle) {
			return gateway.Token
		}
	}
	return ""
}
