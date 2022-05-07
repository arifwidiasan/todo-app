package main

import (
	conf "github.com/arifwidiasan/todo-app/config"
	rest "github.com/arifwidiasan/todo-app/handler/restecho"

	"github.com/labstack/echo/v4"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()

	rest.RegisterUserGroupAPI(e, config)

	e.Logger.Fatal(e.Start((config.SERVER_ADDRESS)))
}
