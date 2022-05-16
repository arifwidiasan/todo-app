package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	//"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// CreateActivity godoc
// @Summary Create New Activity.
// @Description create new activity for user.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity	body	model.CreateActivity	true	"JSON"
// @Success	201	{object} model.CreateActivity
// @Failure 400 {object} model.JWTNotFound
// @Failure 404 {object} model.UserNotFound
// @Failure 500 {object} model.FailCreateActivity
// @Router /{username}/activities [POST]
func (ce *EchoController) CreateActivityController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity := model.Activity{}
	c.Bind(&activity)
	activity = ce.svc.ReplaceActivityName(int(users.ID), activity)

	err = ce.svc.CreateActivityService(activity)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	latest_activity, _ := ce.svc.GetLatestActivityService()

	access := model.Access{}
	err = ce.svc.CreateAccessService(users.ID, latest_activity.ID, true, access)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":      "success",
		"name activity": activity.Activity_Name,
	})
}

// GetAllActivity godoc
// @Summary Get All Activity.
// @Description get all activity by user.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Success	200	{object} model.Activity
// @Failure 400 {object} model.JWTNotFound
// @Failure 404 {object} model.UserNotFound
// @Router /{username}/activities [GET]
func (ce *EchoController) GetAllActivityController(c echo.Context) error {
	username := c.Param("username")
	_, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity := ce.svc.GetAllActivityService(username)

	return c.JSON(200, map[string]interface{}{
		"messages":      "success",
		"List Activity": activity,
	})
}

// GetAllArchivedActivity godoc
// @Summary Get All Archived Activity.
// @Description get all archived activity by user.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Success	200	{object} model.Activity
// @Failure 400 {object} model.JWTNotFound
// @Failure 404 {object} model.UserNotFound
// @Router /{username}/activities/archives [GET]
func (ce *EchoController) GetAllArchiveActivityController(c echo.Context) error {
	username := c.Param("username")
	_, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity := ce.svc.GetAllArchiveActivityService(username)

	return c.JSON(200, map[string]interface{}{
		"messages":      "success",
		"List Activity": activity,
	})
}

// GetActivity godoc
// @Summary Get a Activity.
// @Description get a activity by activity name.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{object} model.Activity
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Router /{username}/activities/{activity_name} [GET]
func (ce *EchoController) GetActivityController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	name := c.Param("activity_name")

	res, err := ce.svc.GetActivityByNameService(name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessService(users.ID, res.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"activity": res,
	})
}

// UpdateActivity godoc
// @Summary Update a Activity.
// @Description update a activity by activity name.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	activity	body	model.CreateActivity	true	"JSON"
// @Success	200	{object} model.ActivityUpdated
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.FailUpdateActivity
// @Router /{username}/activities/{activity_name} [PUT]
func (ce *EchoController) UpdateActivityController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	name := c.Param("activity_name")

	res, err := ce.svc.GetActivityByNameService(name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessOwnerService(users.ID, res.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity",
		})
	}

	id := res.ID

	activity := model.Activity{}
	c.Bind(&activity)
	activity = ce.svc.ReplaceActivityName(int(users.ID), activity)

	err = ce.svc.UpdateActivityService(int(id), activity)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "edited",
		"activity": activity.Activity_Name,
	})
}

// DeleteActivity godoc
// @Summary Delete a Activity.
// @Description Delete a activity by activity name.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{object} model.ActivityDeleted
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccessOwner
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Router /{username}/activities/{activity_name} [DELETE]
func (ce *EchoController) DeleteActivityController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	name := c.Param("activity_name")

	res, err := ce.svc.GetActivityByNameService(name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessOwnerService(users.ID, res.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity or not owner",
		})
	}

	id := res.ID
	ce.svc.DeleteAllAccessService(int(id))
	ce.svc.DeleteAllTaskService(int(id))

	err = ce.svc.DeleteActivityByIDService(int(id))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no activity found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "activity deleted",
	})
}

// ArchiveActivity godoc
// @Summary Archive a Activity.
// @Description archive a activity by activity name.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{object} model.ActivityArchived
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccessOwner
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.FailUpdateActivity
// @Router /{username}/activities/{activity_name}/archive [PUT]
func (ce *EchoController) ArchiveActivityController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	name := c.Param("activity_name")

	res, err := ce.svc.GetActivityByNameService(name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessOwnerService(users.ID, res.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity",
		})
	}

	res.Archive()

	err = ce.svc.ArchiveActivityService(int(res.ID), res)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "activity archived",
		"activity": res,
	})
}

// RestoreActivity godoc
// @Summary Restore a Activity.
// @Description restore a activity by activity name.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{object} model.ActivityRestored
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccessOwner
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.FailUpdateActivity
// @Router /{username}/activities/{activity_name}/archive [DELETE]
func (ce *EchoController) RestoreActivityController(c echo.Context) error {
	username := c.Param("username")
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	name := c.Param("activity_name")

	res, err := ce.svc.GetActivityByNameService(name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessOwnerService(users.ID, res.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity",
		})
	}

	res.Restore()

	err = ce.svc.ArchiveActivityService(int(res.ID), res)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "activity is restored",
		"activity": res,
	})
}
