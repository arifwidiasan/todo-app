package handler

import (
	"github.com/arifwidiasan/todo-app/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateActivityController(c echo.Context) error {
	activity := model.Activity{}
	c.Bind(&activity)

	err := ce.svc.CreateActivityService(activity)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages": "success",
		"users":    activity,
	})
}
