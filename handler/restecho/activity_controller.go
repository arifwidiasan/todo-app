package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateActivityController(c echo.Context) error {
	username := c.Param("username")
	users, _ := ce.svc.GetUserByUsernameService(username)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	err := ce.svc.CheckAuth(int(users.ID), int(claim["id"].(float64)))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
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
	err = ce.svc.CreateAccessService(users.ID, latest_activity.ID, access)
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
	users, _ := ce.svc.GetUserByUsernameService(username)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)
	err := ce.svc.CheckAuth(int(users.ID), int(claim["id"].(float64)))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "unauthorized",
		})
	}

	activity := ce.svc.GetAllActivityService(username)

	return c.JSON(200, map[string]interface{}{
		"messages":      "success",
		"List Activity": activity,
	})
}
