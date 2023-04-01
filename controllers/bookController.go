package controllers

import (
	"golang_microservices/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var bookDatas = []models.Book{}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": models.GetBooks(),
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.CreateBook(newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}

func GetBookById(ctx *gin.Context) {
	bookId := ctx.Param("id")

	book, err := models.GetBook(bookId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Book Not Found",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"book": book,
		})
	}

}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	var updatedBook models.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookUpdated, err := models.UpdateBook(bookId, updatedBook)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"book": bookUpdated,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")

	count, _ := models.DeleteBook(bookId)

	if count == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Book deleted successfully",
		})

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Deleted",
		})
	}
}
