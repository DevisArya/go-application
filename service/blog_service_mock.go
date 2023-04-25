package service

import (
	"errors"
	"praktikum/models"

	"github.com/stretchr/testify/mock"
)

type BlogRepositoryMock struct {
	Mock mock.Mock
}

var dataBlog = []models.Blog{
	{
		ID:      1,
		UserId:  1,
		Title:   "Title test",
		Content: "Content test",
	},

	{
		ID:      2,
		UserId:  2,
		Title:   "Title test 2",
		Content: "Content test 2",
	},
}

func (um *BlogRepositoryMock) CreateBlogController(blog *models.Blog) error {

	if blog.UserId == 0 || blog.Title == "" || blog.Content == "" {
		return errors.New("please complete all fields")
	} else {
		return nil
	}
}
func (um *BlogRepositoryMock) GetBlogController(id int) (error, interface{}) {
	found := 0
	for _, val := range dataBlog {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("Blog not found"), nil
	}

	return nil, dataBlog[id-1]
}

func (um *BlogRepositoryMock) GetBlogsController() (error, interface{}) {
	if dataBlog == nil {
		return errors.New("empty data"), nil
	} else {
		return nil, dataBlog
	}
}

func (um *BlogRepositoryMock) DeleteBlogController(id int) error {
	found := 0
	for _, val := range dataBlog {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("Blog not found")
	}

	return nil
}
func (um *BlogRepositoryMock) UpdateBlogController(dataUpdate *models.Blog, id int) error {

	if dataUpdate.UserId == 0 && dataUpdate.Content == "" && dataUpdate.Title == "" {
		return errors.New("please fill in the update data")
	}
	found := 0
	for _, val := range dataBlog {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("Blog not found")
	}

	return nil
}
