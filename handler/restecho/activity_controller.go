package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// CreateActivity godoc
// @Summary Create New Activity.
// @Description create new activity for user.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	activity	body	docs.CreateActivity	true	"JSON"
// @Success	201	{string} string "created
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "internal server error"
// @Router /activities [POST]
func (ce *EchoController) CreateActivityController(c echo.Context) error {
	username := ce.svc.ClaimToken(c.Get("user").(*jwt.Token))

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
		"activity_name": activity.Activity_Name,
	})
}

// GetAllActivity godoc
// @Summary Get All Activity.
// @Description get all activity by user.
// @Tags Activity
// @Accept json
// @Produce json
// @Success	200	{string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "not found"
// @Router /activities [GET]
func (ce *EchoController) GetAllActivityController(c echo.Context) error {
	username := ce.svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity := ce.svc.GetAllActivityService(username)

	return c.JSON(200, map[string]interface{}{
		"messages":   "success",
		"activities": activity,
	})
}

// GetAllArchivedActivity godoc
// @Summary Get All Archived Activity.
// @Description get all archived activity by user.
// @Tags Activity
// @Accept json
// @Produce json
// @Success	200	{string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 404 {object} string "not found"
// @Router /activities/archives [GET]
func (ce *EchoController) GetAllArchiveActivityController(c echo.Context) error {
	username := ce.svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	activity := ce.svc.GetAllArchiveActivityService(username)

	return c.JSON(200, map[string]interface{}{
		"messages":   "success",
		"activities": activity,
	})
}

// GetActivity godoc
// @Summary Get a Activity.
// @Description get a activity by activity name.
// @Tags Activity
// @Accept json
// @Produce json
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "no access"
// @Failure 404 {string} string "not found"
// @Router /activities/{activity_name} [GET]
func (ce *EchoController) GetActivityController(c echo.Context) error {
	username := ce.svc.ClaimToken(c.Get("user").(*jwt.Token))

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
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	activity	body	docs.CreateActivity	true	"JSON"
// @Success	200	{string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "no access"
// @Failure 404 {string} string "not found"
// @Router /activities/{activity_name} [PUT]
func (ce *EchoController) UpdateActivityController(c echo.Context) error {
	username := ce.svc.ClaimToken(c.Get("user").(*jwt.Token))

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
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "no access"
// @Failure 404 {string} string "not found"
// @Router /activities/{activity_name} [DELETE]
func (ce *EchoController) DeleteActivityController(c echo.Context) error {
	username := ce.svc.ClaimToken(c.Get("user").(*jwt.Token))

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
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "no access"
// @Failure 404 {string} string "not found"
// @Router /activities/{activity_name}/archive [PUT]
func (ce *EchoController) ArchiveActivityController(c echo.Context) error {
	username := ce.svc.ClaimToken(c.Get("user").(*jwt.Token))

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
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "no access"
// @Failure 404 {string} string "not dound"
// @Router /activities/{activity_name}/archive [DELETE]
func (ce *EchoController) RestoreActivityController(c echo.Context) error {
	username := ce.svc.ClaimToken(c.Get("user").(*jwt.Token))

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
