package service

import (
	"errors"
	"praktikum/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

var data = []models.User{
	{
		ID:       1,
		Name:     "devis",
		Email:    "Devis@gmail.com",
		Password: "123",
	},
	{
		ID:       2,
		Name:     "devis2",
		Email:    "Devis2@gmail.com",
		Password: "12345"},
}

func (um *UserRepositoryMock) CreateUserController(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("please complete all fields")
	} else {
		return nil
	}
}
func (um *UserRepositoryMock) GetUserController(id int) (error, interface{}) {
	found := 0
	for _, val := range data {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found"), nil
	}

	return nil, data[id-1]
}

func (um *UserRepositoryMock) GetUsersController() (error, interface{}) {
	if data == nil {
		return errors.New("empty data"), nil
	} else {
		return nil, data
	}
}

func (um *UserRepositoryMock) DeleteUserController(id int) error {
	found := 0
	for _, val := range data {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found")
	}

	return nil
}
func (um *UserRepositoryMock) UpdateUserController(dataUpdate *models.User, id int) error {

	if dataUpdate.Name == "" && dataUpdate.Email == "" && dataUpdate.Password == "" {
		return errors.New("please fill in the update data")
	}
	found := 0
	for _, val := range data {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found")
	}

	return nil
}
func (um *UserRepositoryMock) LoginUserController(user *models.User) (error, string) {

	if user.Email == "" && user.Password == "" {
		return errors.New("please fill in the form"), ""
	}

	var idx int
	found := 0
	for i, val := range data {
		if val.Email == user.Email {
			idx = i
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found"), ""
	}

	if data[idx].Password != user.Password {
		return errors.New("fail login"), ""
	}

	*user = data[idx]

	return nil, ""
}
