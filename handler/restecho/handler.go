package handler

import (
	"github.com/arifwidiasan/todo-app/config"
	"github.com/arifwidiasan/todo-app/database"
	"github.com/arifwidiasan/todo-app/domain"

	m "github.com/arifwidiasan/todo-app/middleware"
	"github.com/arifwidiasan/todo-app/model"
	"github.com/arifwidiasan/todo-app/repository"
	"github.com/arifwidiasan/todo-app/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	_ = NewData()
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

	apiUser := e.Group("/v1") //middleware.Logger(),
	//middleware.CORS(),
	//m.APIKEYMiddleware,

	m.LogMiddleware(e)
	apiUser.POST("/login", cont.LoginUserController)
	apiUser.POST("/register", cont.CreateUserController)
	apiUser.GET("/:username", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("/:username/activities", cont.CreateActivityController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
		},
	))
}

type Datas struct {
}

func (d *Datas) CreateUsers(user model.User) error {
	panic("impl")
}
func (d *Datas) GetUserByUsername(username string) (user model.User, err error) {
	panic("impl")
}
func (d *Datas) GetUserByID(id int) (user model.User, err error) {
	panic("impl")
}

func (d *Datas) CreateActivity(activity model.Activity) error {
	panic("impl")
}
func (d *Datas) GetLatestActivity() (activity model.Activity, err error) {
	panic("impl")
}

func (d *Datas) CreateAccess(access model.Access) error {
	panic("impl")
}

func NewData() domain.AdapterRepository {
	return &Datas{}
}
