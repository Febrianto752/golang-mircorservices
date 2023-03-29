package controllers

import (
	"fmt"
	"golang_microservices/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type Book struct {
// 	Id     int    `json:"book_id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// 	Desc   string `json:"desc"`
// }

var bookDatas = []models.Book{}

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

	// newBook.Id = len(bookDatas) + 1
	// bookDatas = append(bookDatas, newBook)
	models.CreateBook(newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}

func GetBookById(ctx *gin.Context) {
	bookId := ctx.Param("id")
	// bookExist := false
	var bookData models.Book

	bookData = models.GetBook(bookId)

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})

}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	var updatedBook models.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rowUpdated := models.UpdateBook(bookId, updatedBook)

	if rowUpdated == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Updated",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Updated",
		})
	}

}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	bookExist := false
	var bookIndex int

	for i, book := range bookDatas {
		if bookId == strconv.Itoa(book.Id) {
			bookExist = true
			bookIndex = i
			break
		}
	}

	if !bookExist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookId),
		})

		return
	}

	copy(bookDatas[bookIndex:], bookDatas[bookIndex+1:])
	bookDatas[len(bookDatas)-1] = models.Book{}
	bookDatas = bookDatas[:len(bookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deleted", bookId),
	})

}
