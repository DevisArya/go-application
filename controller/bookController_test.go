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

func TestCreateBookControllerValid(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	var dataBook = models.Book{
		ID:        1,
		Title:     "Title test",
		Author:    "Author test",
		Publisher: "Publisher test",
	}

	bookRepository.Mock.On("CreateBookController", &dataBook).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(dataBook)
	req := httptest.NewRequest(http.MethodPost, "/books/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateBookController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestCreateBookControllerInvalid(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	data := models.Book{
		ID:    1,
		Title: "test",
	}
	bookRepository.Mock.On("CreateBookController", &data).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/books/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateBookController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetBookControllerValid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("GetBookController", 1).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	GetBookController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestGetBookControllerInvalid(t *testing.T) {

	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("GetBookController", 3).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	GetBookController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetBooksController(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("GetBooksController").Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetBooksController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteBookControllerValid(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("DeleteBookController", 1).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteBookController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestDeleteBookControllerInvalid(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	bookRepository.Mock.On("DeleteBookController", 3).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	DeleteBookController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestUpdateBookControllerValid(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	data := models.Book{
		Title: "Test Update",
	}

	bookRepository.Mock.On("UpdateBookController", &data, 1).Return(nil)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/books/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateBookController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestUpdateBookControllerInvalid(t *testing.T) {
	bookRepository := &service.BookRepositoryMock{Mock: mock.Mock{}}
	service.SetBookRepository(bookRepository)

	data := models.Book{}

	bookRepository.Mock.On("UpdateBookController", &data, 1).Return(nil)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/books/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateBookController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
