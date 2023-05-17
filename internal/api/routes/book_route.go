package routes

import (
	"github.com/gin-gonic/gin"
	"modelo-rest-go/internal/api/book/controller"
	"modelo-rest-go/internal/api/book/repository"
	"modelo-rest-go/internal/api/book/service"
)

func BookRoute(s *gin.Engine) {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	controller := controller.NewController(service)
	s.GET("/book", controller.GetAll)
	s.POST("/book", controller.CreateBook)
	s.GET("/book/:bookId", controller.GetBookById)
	s.PUT("/book/:bookId", controller.EditBook)
	s.DELETE("/book/:bookId", controller.DeleteBook)
}
