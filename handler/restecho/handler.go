package handler

import (
	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/database"

	m "github.com/arifwidiasan/todo-app/middleware"
	"github.com/arifwidiasan/todo-app/repository"
	"github.com/arifwidiasan/todo-app/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewServiceUser(repo, conf)

	cont := EchoController{
		svc: svc,
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	apiUser := e.Group("/v1")

	m.LogMiddleware(e)
	apiUser.POST("/login", cont.LoginUserController)

	apiUser.POST("/register", cont.CreateUserController)

	apiUser.GET("", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.GET("/", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.GET("/activities", cont.GetAllActivityController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("/activities", cont.CreateActivityController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.GET("/activities/archives", cont.GetAllArchiveActivityController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.GET("/activities/:activity_name", cont.GetActivityController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/activities/:activity_name", cont.UpdateActivityController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/activities/:activity_name", cont.DeleteActivityController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.GET("/activities/:activity_name/manage", cont.GetAccessUserActivityController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("/activities/:activity_name/manage", cont.AddAccessActivityUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/activities/:activity_name/manage", cont.DeleteOneAccessController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.DELETE("/activities/:activity_name/remove", cont.DeleteOneNonOwnerAccessController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.PUT("/activities/:activity_name/archive", cont.ArchiveActivityController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/activities/:activity_name/archive", cont.RestoreActivityController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.GET("/activities/:activity_name/tasks", cont.GetAllTaskController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("/activities/:activity_name/tasks", cont.CreateTaskController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.GET("/activities/:activity_name/tasks/:id", cont.GetOneTaskController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/activities/:activity_name/tasks/:id", cont.UpdateTaskController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/activities/:activity_name/tasks/:id", cont.DeleteTaskController, middleware.JWT([]byte(conf.JWT_KEY)))

	apiUser.PUT("/activities/:activity_name/tasks/:id/complete", cont.CompleteTaskController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/activities/:activity_name/tasks/:id/complete", cont.UndoCompletedTaskController, middleware.JWT([]byte(conf.JWT_KEY)))
}
