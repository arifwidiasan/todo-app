package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateTaskController(c echo.Context) error {
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

	activity_name := c.Param("activity_name")
	activity, err := ce.svc.GetActivityByNameService(activity_name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessService(users.ID, activity.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity",
		})
	}

	newTask := model.Task{}
	c.Bind(&newTask)
	err = ce.svc.CreateTaskService(activity.ID, newTask)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(201, map[string]interface{}{
		"messages":      "success add task to this activity",
		"name activity": activity.Activity_Name,
		"Task":          newTask,
	})
}

func (ce *EchoController) GetAllTaskController(c echo.Context) error {
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

	activity_name := c.Param("activity_name")
	activity, err := ce.svc.GetActivityByNameService(activity_name)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "activity not found",
		})
	}

	err = ce.svc.CheckAcccessService(users.ID, activity.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity",
		})
	}

	tasks := ce.svc.GetAllTaskService(int(activity.ID))
	return c.JSON(201, map[string]interface{}{
		"messages": "success",
		"Task":     tasks,
	})
}
