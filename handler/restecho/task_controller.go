package handler

import (
	"strconv"

	"github.com/arifwidiasan/todo-app/model"
	//"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// AddTask godoc
// @Summary Create Task of an Activity.
// @Description create task of an activity.
// @Tags Task
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	task	body	model.CreateTask	true	"JSON"
// @Success	201	{object} model.SuccessCreateTask
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 500 {object} model.FailCreateTask
// @Router /{username}/activities/{activity_name}/tasks [POST]
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

// GetAllTask godoc
// @Summary Get All Task in an Activity.
// @Description get all task in an activity.
// @Tags Task
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Success	200	{object} model.Task
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Router /{username}/activities/{activity_name}/tasks [GET]
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

// GetTask godoc
// @Summary Get a Task in an Activity.
// @Description get a task in an activity.
// @Tags Task
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	id	path	string	true	"ID task"
// @Success	200	{object} model.Task
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.TaskNotFound
// @Router /{username}/activities/{activity_name}/tasks/{id} [GET]
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
		return c.JSON(404, map[string]interface{}{
			"messages": "task not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"Task":     task,
	})
}

// UpdateTask godoc
// @Summary Update a Task in an Activity.
// @Description update a task in an activity.
// @Tags Task
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	id	path	string	true	"ID task"
// @Param	task	body	model.CreateTask	true	"JSON"
// @Success	200	{object} model.TaskUpdated
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.TaskNotFound
// @Failure 404 {object} model.FailUpdateTask
// @Router /{username}/activities/{activity_name}/tasks/{id} [PUT]
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
		return c.JSON(404, map[string]interface{}{
			"messages": "task not found or no change",
		})
	}

	result, _ := ce.svc.GetTaskByIDService(task_id_int)
	return c.JSON(200, map[string]interface{}{
		"messages": "success update task",
		"Task":     result,
	})
}

// DeleteTask godoc
// @Summary Delete a Task in an Activity.
// @Description delete a task in an activity.
// @Tags Task
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	id	path	string	true	"ID task"
// @Success	200	{object} model.TaskDeleted
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.TaskNotFound
// @Failure 404 {object} model.FailUpdateTask
// @Router /{username}/activities/{activity_name}/tasks/{id} [DELETE]
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

// CompleteTask godoc
// @Summary Complete a Task in an Activity.
// @Description complete a task in an activity.
// @Tags Task
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	id	path	string	true	"ID task"
// @Success	200	{object} model.TaskCompleted
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.TaskNotFound
// @Router /{username}/activities/{activity_name}/tasks/{id}/complete [PUT]
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

// UndoCompleteTask godoc
// @Summary Undo a Completed Task in an Activity.
// @Description undo a completed task in an activity.
// @Tags Task
// @Accept json
// @Produce json
// @Param	username	path	string	true	"Username"
// @Param	activity_name	path	string	true	"Activity Name"
// @Param	id	path	string	true	"ID task"
// @Success	200	{object} model.TaskUndo
// @Failure 400 {object} model.JWTNotFound
// @Failure 401 {object} model.NoAccess
// @Failure 404 {object} model.UserNotFound
// @Failure 404 {object} model.ActivityNotFound
// @Failure 404 {object} model.TaskNotFound
// @Router /{username}/activities/{activity_name}/tasks/{id}/complete [DELETE]
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
