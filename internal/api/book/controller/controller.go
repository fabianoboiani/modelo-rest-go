package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"modelo-rest-go/internal/api/book/service"
	"modelo-rest-go/internal/api/responses"
	"modelo-rest-go/internal/pkg/models"
	"net/http"
)

type controller struct {
	service service.Service
}

var validate = validator.New()

func NewController(service service.Service) *controller {
	return &controller{service: service}
}

func (c *controller) GetAll(ctx *gin.Context) {

	books, err := c.service.GetAll()

	if err != nil {
		ctx.JSON(err.Status, responses.Response{Status: err.Status, Message: "error to get all", Data: gin.H{"error": err.Error}})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: gin.H{"books": books}})
	return
}

func (c *controller) CreateBook(ctx *gin.Context) {

	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "Error", Data: gin.H{"Erro": err.Error()}})
		return
	}

	if validateErr := validate.Struct(&book); validateErr != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusInternalServerError, Message: "Error", Data: gin.H{"Erro": validateErr.Error()}})
		return
	}

	result, err := c.service.CreateBook(&book)

	if err != nil {
		ctx.JSON(err.Status, responses.Response{Status: err.Status, Message: "Error", Data: gin.H{"Erro": err.Error}})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: gin.H{"result": result}})
}

func (c *controller) GetBookById(ctx *gin.Context) {

	bookId := ctx.Param("bookId")

	book, err := c.service.GetBookById(bookId)

	if err != nil {
		ctx.JSON(err.Status, responses.Response{Status: err.Status, Message: "error", Data: gin.H{"error": err.Error}})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: gin.H{"book": book}})
	return
}

func (c *controller) EditBook(ctx *gin.Context) {

	bookId := ctx.Param("bookId")
	var book models.Book

	//validate the request body
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error on data", Data: gin.H{"error": err.Error()}})
		return
	}
	objId, _ := primitive.ObjectIDFromHex(bookId)
	book.Id = objId

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&book); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: gin.H{"error": validationErr.Error()}})
		return
	}

	updatedBook, err := c.service.EditBook(book)

	if err != nil {
		ctx.JSON((*err).Status, responses.Response{Status: (*err).Status, Message: "error", Data: gin.H{"error": (*err).Error}})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: gin.H{"book": *updatedBook}})
	return
}

func (c *controller) DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	err := c.service.DeleteBook(bookId)

	if err != nil {
		if err.Status == http.StatusNotFound {
			ctx.JSON(err.Status, responses.Response{Status: err.Status, Message: "Book id not found", Data: gin.H{"error": (*err).Error}})
			return
		} else {
			ctx.JSON(err.Status, responses.Response{Status: err.Status, Message: "error", Data: gin.H{"error": (*err).Error}})
			return
		}
	}

	ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: gin.H{}})
	return
}
