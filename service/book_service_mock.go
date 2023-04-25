package service

import (
	"errors"
	"praktikum/models"

	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	Mock mock.Mock
}

var dataBook = []models.Book{
	{
		ID:        1,
		Title:     "Title test",
		Author:    "Author test",
		Publisher: "Publisher test",
	},
	{
		ID:        2,
		Title:     "Title test 2",
		Author:    "Author test 2",
		Publisher: "Publisher test 2",
	},
}

func (um *BookRepositoryMock) CreateBookController(book *models.Book) error {
	if book.Title == "" || book.Author == "" || book.Publisher == "" {
		return errors.New("please complete all fields")
	} else {
		return nil
	}
}
func (um *BookRepositoryMock) GetBookController(id int) (error, interface{}) {
	found := 0
	for _, val := range dataBook {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("Book not found"), nil
	}

	return nil, dataBook[id-1]
}

func (um *BookRepositoryMock) GetBooksController() (error, interface{}) {
	if dataBook == nil {
		return errors.New("empty data"), nil
	} else {
		return nil, dataBook
	}
}

func (um *BookRepositoryMock) DeleteBookController(id int) error {
	found := 0
	for _, val := range dataBook {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("Book not found")
	}

	return nil
}
func (um *BookRepositoryMock) UpdateBookController(dataUpdate *models.Book, id int) error {

	if dataUpdate.Author == "" && dataUpdate.Publisher == "" && dataUpdate.Title == "" {
		return errors.New("please fill in the update data")
	}
	found := 0
	for _, val := range dataBook {
		if val.ID == id {
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("Book not found")
	}

	return nil
}
