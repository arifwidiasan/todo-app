package handler

import (
	"net/http"

	"github.com/arifwidiasan/todo-app/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := ce.svc.CreateUserService(user)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages": "success",
		"users":    user,
	})
}

func (ce *EchoController) GetUserController(c echo.Context) error {
	username := c.Param("username")

	res, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})
}

func (ce *EchoController) LoginUserController(c echo.Context) error {
	userLogin := model.LoginRequest{}

	c.Bind(&userLogin)

	token, statusCode := ce.svc.LoginUser(userLogin.Username, userLogin.Password)
	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
			"messages": "username atau password salah",
		}, "  ")

	case http.StatusInternalServerError:
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal",
		}, "  ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token":    token,
	}, "  ")
}
