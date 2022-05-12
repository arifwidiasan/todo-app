package handler

import (
	"strconv"

	"github.com/arifwidiasan/todo-app/model"
	//"github.com/golang-jwt/jwt"
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
	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"Task":     tasks,
	})
}

func (ce *EchoController) GetOneTaskController(c echo.Context) error {
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

	err = ce.svc.CheckAcccessService(users.ID, activity.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity",
		})
	}

	task_id := c.Param("id")
	task_id_int, _ := strconv.Atoi(task_id)
	task, err := ce.svc.GetTaskByIDService(task_id_int)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "task not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"Task":     task,
	})
}

func (ce *EchoController) UpdateTaskController(c echo.Context) error {
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

	err = ce.svc.CheckAcccessService(users.ID, activity.ID)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "you dont have access to this activity",
		})
	}

	task_id := c.Param("id")
	task_id_int, _ := strconv.Atoi(task_id)

	task := model.Task{}
	c.Bind(&task)
	err = ce.svc.UpdateTaskService(task_id_int, task)
	if err != nil {
		return c.JSON(401, map[string]interface{}{
			"messages": "task not found or no change",
		})
	}

	result, _ := ce.svc.GetTaskByIDService(task_id_int)
	return c.JSON(200, map[string]interface{}{
		"messages": "success update task",
		"Task":     result,
	})
}

func (ce *EchoController) DeleteTaskController(c echo.Context) error {
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
			"messages": "you dont have access to this activity or not owner",
		})
	}

	task_id := c.Param("id")
	task_id_int, _ := strconv.Atoi(task_id)

	err = ce.svc.DeleteTaskByIDService(task_id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no task found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "task deleted",
	})
}

func (ce *EchoController) CompleteTaskController(c echo.Context) error {
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

	task_id := c.Param("id")
	task_id_int, _ := strconv.Atoi(task_id)

	result, err := ce.svc.GetTaskByIDService(task_id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no task found",
		})
	}

	result.Complete()

	err = ce.svc.CompleteTaskService(int(result.ID), result)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "task completed",
		"activity": result,
	})
}

func (ce *EchoController) UndoCompletedTaskController(c echo.Context) error {
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

	task_id := c.Param("id")
	task_id_int, _ := strconv.Atoi(task_id)

	result, err := ce.svc.GetTaskByIDService(task_id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no task found",
		})
	}

	result.Undo()

	err = ce.svc.CompleteTaskService(int(result.ID), result)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success undo completed task",
		"activity": result,
	})
}
