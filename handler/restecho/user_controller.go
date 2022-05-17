package handler

import (
	"net/http"

	"github.com/arifwidiasan/todo-app/model"
	"github.com/labstack/echo/v4"
)

// CreateUser godoc
// @Summary Create/Register New User.
// @Description create new user with username, user_email. user_pass.
// @Tags User
// @Accept json
// @Produce json
// @Param	user	body	model.CreateUser	true	"JSON username, user_email, and user_pass"
// @Success	201	{object} model.CreateUser
// @Failure 500 {object} model.FailCreateUser "Internal Server Error"
// @Router /register [POST]
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

// GetUser godoc
// @Summary Get a User.
// @Description get a user information by username.
// @Tags User
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Success	200	{object} model.UserFound "success"
// @Failure 400 {object} model.JWTNotFound
// @Failure 404 {object} model.UserNotFound
// @Router /{username} [GET]
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

// LoginUser godoc
// @Summary Login User.
// @Description login user to get jwt token.
// @Tags User
// @Accept json
// @Produce json
// @Param	user	body	model.LoginRequest	true	"JSON username and user_pass"
// @Success	200	{object} model.SuccessLoginUser
// @Failure 401 {object} model.FailCreateUser "unauthorized"
// @Failure 500 {object} model.FailCreateUser "internal server error"
// @Router /login [POST]
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
