package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	//"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

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
