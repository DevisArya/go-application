package service

import (
	"errors"
	"net/http"
	"praktikum/config"
	"praktikum/models"

	"github.com/labstack/echo/v4"
)

type IBlogService interface {
	CreateBlogController(blog *models.Blog) error
	GetBlogController(id int) (error, interface{})
	GetBlogsController() (error, interface{})
	DeleteBlogController(id int) error
	UpdateBlogController(dataUpdate *models.Blog, id int) error
}

type BlogRepository struct {
	Func IBlogService
}

var blogRepository IBlogService

func init() {
	bg := &BlogRepository{}
	bg.Func = bg

	blogRepository = bg
}
func GetBlogRepository() IBlogService {
	return blogRepository
}
func SetBlogRepository(ur IBlogService) {
	blogRepository = ur
}

func (u *BlogRepository) CreateBlogController(blog *models.Blog) error {
	if blog.UserId == 0 || blog.Title == "" || blog.Content == "" {
		return errors.New("please complete all fields")
	}
	if err := config.DB.Save(&blog); err != nil {
		return err.Error
	}
	return nil
}

func (u *BlogRepository) GetBlogController(id int) (err error, res interface{}) {
	var blog models.Blog
	if result := config.DB.Where("id = ?", id).First(&blog); result.Error != nil {
		return errors.New("Blog not found"), nil
	}
	return nil, blog
}

func (u *BlogRepository) GetBlogsController() (err error, res interface{}) {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), nil
	}
	return nil, blogs
}

func (u *BlogRepository) DeleteBlogController(id int) error {

	result := config.DB.Delete(&models.Blog{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "Blog not found",
		})
	}
	return nil
}

func (u *BlogRepository) UpdateBlogController(dataUpdate *models.Blog, id int) error {
	var blogs models.Blog
	result := config.DB.Model(&blogs).Where("id = ?", id).Updates(&dataUpdate)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "Blog not found",
		})
	}
	return nil
}
