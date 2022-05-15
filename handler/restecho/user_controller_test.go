package handler

import (
	//"errors"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserControllerValid(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateUserService", mock.Anything).Return(nil).Once()

	usrController := EchoController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("success", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserController(echoContext)
		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestCreateUserControllerNotValid(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateUserService", mock.Anything).Return(errors.New("error")).Once()

	usrController := EchoController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("fail", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserController(echoContext)
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetUserControllerValid(t *testing.T) {
	svc := MockSvc{}

	svc.On("GetUserByUsernameService", mock.Anything).Return(nil).Once()

	usrController := EchoController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("success", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/:username", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetUserController(echoContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetUserControllerNotValid(t *testing.T) {
	svc := MockSvc{}

	svc.On("GetUserByUsernameService", mock.Anything).Return(errors.New("new")).Once()

	usrController := EchoController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("fail", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/:username", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetUserController(echoContext)
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestLoginUserControllerValid(t *testing.T) {
	svc := MockSvc{}

	svc.On("LoginUser", mock.Anything).Return(nil).Once()

	usrController := EchoController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("success", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.LoginUserController(echoContext)
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}
