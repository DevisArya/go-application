package service

import (
	"errors"
	"net/http"
	"praktikum/config"
	md "praktikum/middleware"
	"praktikum/models"

	"github.com/labstack/echo/v4"
)

type IUserService interface {
	CreateUserController(user *models.User) error
	GetUserController(id int) (error, interface{})
	GetUsersController() (error, interface{})
	DeleteUserController(id int) error
	UpdateUserController(dataUpdate *models.User, id int) error
	LoginUserController(user *models.User) (error, string)
}

type UserRepository struct {
	Func IUserService
}

var userRepository IUserService

func init() {
	ur := &UserRepository{}
	ur.Func = ur

	userRepository = ur
}
func GetUserRepository() IUserService {
	return userRepository
}
func SetUserRepository(ur IUserService) {
	userRepository = ur
}

func (u *UserRepository) CreateUserController(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("please complete all fields")
	}
	if err := config.DB.Save(&user); err != nil {
		return err.Error
	}
	return nil
}

func (u *UserRepository) GetUserController(id int) (err error, res interface{}) {
	var user models.User
	if result := config.DB.Where("id = ?", id).First(&user); result.Error != nil {
		return errors.New("user not found"), nil
	}
	return nil, user
}

func (u *UserRepository) GetUsersController() (err error, res interface{}) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), nil
	}
	return nil, users
}

func (u *UserRepository) DeleteUserController(id int) error {

	result := config.DB.Delete(&models.User{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}
	return nil
}

func (u *UserRepository) UpdateUserController(dataUpdate *models.User, id int) error {
	var users models.User
	result := config.DB.Model(&users).Where("id = ?", id).Updates(&dataUpdate)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}
	return nil
}

func (u *UserRepository) LoginUserController(user *models.User) (error, string) {

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "fail login",
		}), ""
	}

	token, err := md.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		}), ""
	}
	return nil, token
}
