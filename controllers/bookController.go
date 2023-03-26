package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int
	Title  string
	Author string
	Desc   string
}

var bookDatas = []Book{}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": bookDatas,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.Id = len(bookDatas) + 1
	bookDatas = append(bookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
