package main

import (
	conf "github.com/arifwidiasan/todo-app/config"
	_ "github.com/arifwidiasan/todo-app/docs"
	rest "github.com/arifwidiasan/todo-app/handler/restecho"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

// @title To-do List App API
// @version 2.0
// @description Documentation of To-do List App API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
// @schemes http
func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	rest.RegisterUserGroupAPI(e, config)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start((config.SERVER_ADDRESS)))
}
