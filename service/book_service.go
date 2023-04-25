package service

import (
	"errors"
	"net/http"
	"praktikum/config"
	"praktikum/models"

	"github.com/labstack/echo/v4"
)

type IBookService interface {
	CreateBookController(book *models.Book) error
	GetBookController(id int) (error, interface{})
	GetBooksController() (error, interface{})
	DeleteBookController(id int) error
	UpdateBookController(dataUpdate *models.Book, id int) error
}

type BookRepository struct {
	Func IBookService
}

var bookRepository IBookService

func init() {
	br := &BookRepository{}
	br.Func = br

	bookRepository = br
}
func GetBookRepository() IBookService {
	return bookRepository
}
func SetBookRepository(ur IBookService) {
	bookRepository = ur
}

func (u *BookRepository) CreateBookController(book *models.Book) error {
	if book.Title == "" || book.Author == "" || book.Publisher == "" {
		return errors.New("please complete all fields")
	}
	if err := config.DB.Save(&book); err != nil {
		return err.Error
	}
	return nil
}

func (u *BookRepository) GetBookController(id int) (err error, res interface{}) {
	var book models.Book
	if result := config.DB.Where("id = ?", id).First(&book); result.Error != nil {
		return errors.New("Book not found"), nil
	}
	return nil, book
}

func (u *BookRepository) GetBooksController() (err error, res interface{}) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), nil
	}
	return nil, books
}

func (u *BookRepository) DeleteBookController(id int) error {

	result := config.DB.Delete(&models.Book{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "Book not found",
		})
	}
	return nil
}

func (u *BookRepository) UpdateBookController(dataUpdate *models.Book, id int) error {
	var books models.Book
	result := config.DB.Model(&books).Where("id = ?", id).Updates(&dataUpdate)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "Book not found",
		})
	}
	return nil
}
