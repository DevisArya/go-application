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

func TestCreateBlogControllerValid(t *testing.T) {
	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	var dataBlog = models.Blog{

		ID:      1,
		UserId:  1,
		Title:   "Title test",
		Content: "Content test",
	}

	blogRepository.Mock.On("CreateBlogController", &dataBlog).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(dataBlog)
	req := httptest.NewRequest(http.MethodPost, "/blogs/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateBlogController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestCreateBlogControllerInvalid(t *testing.T) {
	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	data := models.Blog{
		Title: "test",
	}
	blogRepository.Mock.On("CreateBlogController", &data).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/blogs/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateBlogController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetBlogControllerValid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("GetBlogController", 1).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	GetBlogController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestGetBlogControllerInvalid(t *testing.T) {

	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("GetBlogController", 3).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	GetBlogController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetBlogsController(t *testing.T) {
	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("GetBlogsController").Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetBlogsController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteBlogControllerValid(t *testing.T) {
	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("DeleteBlogController", 1).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteBlogController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestDeleteBlogControllerInvalid(t *testing.T) {
	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	blogRepository.Mock.On("DeleteBlogController", 3).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/blogs/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	DeleteBlogController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
func TestUpdateBlogControllerValid(t *testing.T) {
	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	data := models.Blog{
		Title: "Test Update",
	}

	blogRepository.Mock.On("UpdateBlogController", &data, 1).Return(nil)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/blogs/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateBlogController(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestUpdateBlogControllerInvalid(t *testing.T) {
	blogRepository := &service.BlogRepositoryMock{Mock: mock.Mock{}}
	service.SetBlogRepository(blogRepository)

	data := models.Blog{}

	blogRepository.Mock.On("UpdateBlogController", &data, 1).Return(nil)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/blogs/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateBlogController(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
