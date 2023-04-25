package controller

import (
	"net/http"
	m "praktikum/models"
	"praktikum/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all blogs
func GetBlogsController(c echo.Context) error {

	err, res := service.GetBlogRepository().GetBlogsController()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all Blogs",
		"blogs":   res,
	})
}

// get Blog by id
func GetBlogController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	err, res := service.GetBlogRepository().GetBlogController(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get blog",
		"blog":     res,
	})
}

// create new Blog
func CreateBlogController(c echo.Context) error {
	blog := m.Blog{}
	c.Bind(&blog)

	if err := service.GetBlogRepository().CreateBlogController(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete Blog by id
func DeleteBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := service.GetBlogRepository().DeleteBlogController(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete blog",
	})
}

// update Blog by id
func UpdateBlogController(c echo.Context) error {
	dataUpdate := m.Blog{}
	c.Bind(&dataUpdate)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := service.GetBlogRepository().UpdateBlogController(&dataUpdate, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Update blog",
	})
}
