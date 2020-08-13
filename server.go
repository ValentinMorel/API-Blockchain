package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	echoInstance := echo.New()

	// Logger for server Log
	echoInstance.Use(middleware.Logger())
	// Recover middleware gives the possibility to recover from a panic err
	echoInstance.Use(middleware.Recover())

	// Cross domain access controls -> secure data transfers
	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.HEAD},
	}))

	echoInstance.GET("/", func(echoContext echo.Context) error {
		return echoContext.String(http.StatusOK, "Hello World \n")
	})

	echoInstance.Logger.Fatal(echoInstance.Start(":8080"))
	//Make a request on the API with :
	// curl -X GET localhost:8080/ -> Response : "Hello World"
}
