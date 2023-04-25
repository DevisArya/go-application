package controller

import (
	"net/http"

	m "praktikum/models"
	"praktikum/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {

	err, res := service.GetUserRepository().GetUsersController()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   res,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	err, res := service.GetUserRepository().GetUserController(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"user":     res,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := service.GetUserRepository().DeleteUserController(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := m.User{}
	c.Bind(&user)

	if err := service.GetUserRepository().CreateUserController(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {

	dataUpdate := m.User{}
	c.Bind(&dataUpdate)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := service.GetUserRepository().UpdateUserController(&dataUpdate, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Update user",
	})
}

func LoginUserController(c echo.Context) error {
	user := m.User{}
	c.Bind(&user)

	err, token := service.GetUserRepository().LoginUserController(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	userResponse := m.UserResponse{ID: int(user.ID), Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Login",
		"user":    userResponse,
	})
}
