package controller

import (
	"net/http"
	m "praktikum/models"
	"praktikum/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {

	err, res := service.GetBookRepository().GetBooksController()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   res,
	})
}

// get book by id
func GetBookController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	err, res := service.GetBookRepository().GetBookController(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get book",
		"book":     res,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := m.Book{}
	c.Bind(&book)

	if err := service.GetBookRepository().CreateBookController(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := service.GetBookRepository().DeleteBookController(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	dataUpdate := m.Book{}
	c.Bind(&dataUpdate)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := service.GetBookRepository().UpdateBookController(&dataUpdate, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Update book",
	})
}
