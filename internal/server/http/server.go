package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"net/http"
	"os"
	"shop/internal/server/api"
	"shop/pkg/advancedError"
	"shop/pkg/database"
)

// DEFAULT_PORT defines default application port
const DEFAULT_PORT string = "8000"

// StartHttpServer for running
func StartHttpServer() {
	dbConn := connectToDb()
	router := makeNewApp()

	registerServices(dbConn, router)

	port := appPort()
	if err := router.Start(port); err != nil {
		panic(err)
	}
}

// connectToDb initializes database connection
func connectToDb() *gorm.DB {
	conn, err := database.Conn()
	if err != nil {
		panic(advancedError.New(err, "Starting database connection failed"))
	}
	return conn
}

// makeNewApp and return it
func makeNewApp() *echo.Echo {
	return echo.New()
}

// registerServices in main app
func registerServices(conn *gorm.DB, router *echo.Echo) {
	appendRequiredMiddlewares(router)
	api.AttachAddressHandlerWithAddressDomain(router, conn)
	api.AttachBasketHandlerWithBasketDomain(router, conn)
	api.AttachCommentHandlerWithCommentDomain(router, conn)
	api.AttachContactUsHandlerToContactUsDomain(router, conn)
	api.AttachGatewayHandlerToGatewayDomain(router, conn)
	api.AttachLikeDislikeToItsDomain(router, conn)
	api.AttachPaymentWithRestUrls(router, conn)
	api.AttachPostTypeHandlerToPostTypeDomain(router, conn)
	api.AttachProductHandlerToProductDomain(router, conn)
	api.AttachUserHandlerToUserDomain(router, conn)
}

// appendRequiredMiddlewares to http router function
func appendRequiredMiddlewares(router *echo.Echo) {
	router.Use(middleware.Logger())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	//Todo disable strict routing
}

// appPort gets and returns application port
func appPort() string {
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	return fmt.Sprintf(":%s", port)
}
