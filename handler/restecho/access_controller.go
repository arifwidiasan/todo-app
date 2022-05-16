package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	//"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// AddAccess godoc
// @Summary Add Access User to Activity.
// @Description add access user to activity.
// @Tags Access
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	access	body	model.CreateAccess	true	"JSON"
// @Success	201	{object} model.SuccessCreateAccess
// @Failure 400 {object} model.JWTNotFound
// @Failure 400 {object} model.AddYourself
// @Failure 401 {object} model.NoAccessOwner
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 500 {object} model.FailCreateAccess
// @Router /{username}/activities/{activity_name}/manage [POST]
func (ce *EchoController) AddAccessActivityUserController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity_name := c.Param("activity_name")
	activity, err := ce.svc.GetActivityByNameService(activity_name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessOwnerService(users.ID, activity.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity because you're not the owner",
		})
	}

	newUser := model.User{}
	c.Bind(&newUser)
	if newUser.Username == users.Username {
		return c.JSON(400, map[string]interface{}{
			"messages": "can't add yourself",
		})
	}

	newUser, err = ce.svc.GetUserByUsernameService(newUser.Username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "target username not found",
		})
	}

	access := model.Access{}
	err = ce.svc.CreateAccessService(newUser.ID, activity.ID, false, access)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(201, map[string]interface{}{
		"messages":      "success add another user to this activity",
		"name activity": activity.Activity_Name,
		"username":      newUser.Username,
	})
}

// GetListAccess godoc
// @Summary Get List Access User to Activity.
// @Description get list access user to activity.
// @Tags Access
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{object} model.ListAccess
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccessOwner
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Router /{username}/activities/{activity_name}/manage [GET]
func (ce *EchoController) GetAccessUserActivityController(c echo.Context) error {
	username := c.Param("username")
	users, _ := ce.svc.GetUserByUsernameService(username)
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity_name := c.Param("activity_name")
	activity, err := ce.svc.GetActivityByNameService(activity_name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessOwnerService(users.ID, activity.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity because you're not the owner",
		})
	}

	listAccess := ce.svc.GetAccessUserActivityService(int(activity.ID))

	return c.JSON(200, map[string]interface{}{
		"messages":                   "success",
		"List User in This Activity": listAccess,
	})
}

// DeleteAccess godoc
// @Summary Delete Access Another User to Activity.
// @Description delete access another user to activity.
// @Tags Access
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	access	body	model.CreateAccess	true	"JSON"
// @Success	200	{object} model.SuccessDeleteAccess
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccessOwner
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.FailDeleteAccess
// @Router /{username}/activities/{activity_name}/manage [DELETE]
func (ce *EchoController) DeleteOneAccessController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity_name := c.Param("activity_name")
	activity, err := ce.svc.GetActivityByNameService(activity_name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessOwnerService(users.ID, activity.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity because you're not the owner",
		})
	}

	newUser := model.User{}
	c.Bind(&newUser)
	newUser, err = ce.svc.GetUserByUsernameService(newUser.Username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "target username not found",
		})
	}

	err = ce.svc.DeleteOneAccessService(int(newUser.ID), int(activity.ID))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"messages":      "success delete user to this activity",
		"name activity": activity.Activity_Name,
		"username":      newUser.Username,
	})
}

// RemoveAccess godoc
// @Summary Remove Access User from Activity.
// @Description remove access user from activity.
// @Tags Access
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{object} model.SuccessRemoveAccess
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccessOwner
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.FailDeleteAccess
// @Router /{username}/activities/{activity_name}/remove [DELETE]
func (ce *EchoController) DeleteOneNonOwnerAccessController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity_name := c.Param("activity_name")
	activity, err := ce.svc.GetActivityByNameService(activity_name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.DeleteOneAccessService(int(users.ID), int(activity.ID))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages":      "success delete access, you can't access this activity anymore",
		"name activity": activity.Activity_Name,
	})
}
