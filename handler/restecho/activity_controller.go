package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	"github.com/golang-jwt/jwt"
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

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	err = ce.svc.CheckAuth(int(users.ID), int(claim["id"].(float64)))
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "unauthorized",
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
	users, err := ce.svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "username not found",
		})
	}

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	err = ce.svc.CheckAuth(int(users.ID), int(claim["id"].(float64)))
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "unauthorized",
		})
	}

	activity := ce.svc.GetAllActivityService(username)

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

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	err = ce.svc.CheckAuth(int(users.ID), int(claim["id"].(float64)))
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "unauthorized",
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

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	err = ce.svc.CheckAuth(int(users.ID), int(claim["id"].(float64)))
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "unauthorized",
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

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	err = ce.svc.CheckAuth(int(users.ID), int(claim["id"].(float64)))
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "unauthorized",
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
	err = ce.svc.DeleteAllAccessService(int(id))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no acccess activity found",
		})
	}

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
