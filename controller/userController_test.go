package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"praktikum/models"
	"praktikum/service"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserControllerValid(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	data := models.User{
		ID:       1,
		Name:     "devis",
		Email:    "Devis@gmail.com",
		Password: "123",
	}
	userRepository.Mock.On("CreateUserController", &data).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateUserController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestCreateUserControllerInvalid(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	data := models.User{
		ID:   1,
		Name: "devis",
	}
	userRepository.Mock.On("CreateUserController", &data).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateUserController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetUserControllerValid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUserController", 1).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	GetUserController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestGetUserControllerInvalid(t *testing.T) {

	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUserController", 3).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	GetUserController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetUsersController(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("GetUsersController").Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetUsersController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteUserControllerValid(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("DeleteUserController", 1).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteUserController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestDeleteUserControllerInvalid(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	userRepository.Mock.On("DeleteUserController", 3).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	DeleteUserController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestUpdateUserControllerValid(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	data := models.User{
		Name: "Devis Arya Pratama",
	}

	userRepository.Mock.On("UpdateUserController", &data, 1).Return(nil)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/users/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateUserController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestUpdateUserControllerInvalid(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	data := models.User{}

	userRepository.Mock.On("UpdateUserController", &data, 1).Return(nil)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/users/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateUserController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestLoginUserControllerValid(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	data := models.User{
		Email:    "Devis@gmail.com",
		Password: "123",
	}

	userRepository.Mock.On("LoginUserController", &data).Return(nil)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/login/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	LoginUserController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestLoginUserControllerinvalid(t *testing.T) {
	userRepository := &service.UserRepositoryMock{Mock: mock.Mock{}}
	service.SetUserRepository(userRepository)

	data := models.User{
		Email:    "Devis@gmail.com",
		Password: "1",
	}

	userRepository.Mock.On("LoginUserController", &data).Return(nil)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/login/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	LoginUserController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
