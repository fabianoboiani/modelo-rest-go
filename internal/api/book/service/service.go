package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	repository2 "modelo-rest-go/internal/api/book/repository"
	"modelo-rest-go/internal/pkg/models"
)

type Service interface {
	GetAll() (*[]models.Book, *models.Error)
	CreateBook(book *models.Book) (interface{}, *models.Error)
	GetBookById(bookId string) (*models.Book, *models.Error)
	EditBook(book models.Book) (*models.Book, *models.Error)
	DeleteBook(bookId string) *models.Error
}

type service struct {
	repository repository2.Repository
}

func NewService(repository repository2.Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAll() (*[]models.Book, *models.Error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) CreateBook(book *models.Book) (interface{}, *models.Error) {

	newBook := models.Book{
		Id:       primitive.NewObjectID(),
		Title:    book.Title,
		Category: book.Category,
		Author:   book.Author,
	}

	result, err := s.repository.Insert(&newBook)

	return result, err
}

func (s *service) GetBookById(bookId string) (*models.Book, *models.Error) {
	book, err := s.repository.FindById(bookId)
	return book, err
}

func (s *service) EditBook(book models.Book) (*models.Book, *models.Error) {
	updatedBook, err := s.repository.Update(book)
	return updatedBook, err
}

func (s *service) DeleteBook(bookId string) *models.Error {
	err := s.repository.Delete(bookId)
	return err
}
