package routers

import (
	"golang_microservices/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Logger(),   // untuk log request yang masuk
		gin.Recovery(), // untuk auto restart kalau panic
	)

	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.GetBookById)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
