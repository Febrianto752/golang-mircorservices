package models

import (
	"errors"
	"fmt"
	"golang_microservices/database"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Book struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	BookName  string    `gorm:"not null;unique;type:varchar(100)" json:"book_name"`
	Author    string    `gorm:"not null;type:varchar(100)" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateBook(newBook Book) Book {
	db := database.GetDB()

	book := Book{
		BookName: newBook.BookName,
		Author:   newBook.Author,
	}

	err := db.Create(&book).Error

	if err != nil {
		fmt.Println("Error creating user data :", err)
		panic(err)
	}

	fmt.Println("New Book Data :", book)
	return book
}

func GetBooks() []Book {
	db := database.GetDB()

	var books = []Book{}

	err := db.Find(&books).Error

	if err != nil {
		panic(err)
	}

	return books
}

func GetBook(id string) (Book, error) {
	db := database.GetDB()

	bookId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	var book Book

	err = db.First(&book, "id = ?", bookId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return book, err
		}
		print("Error finding user", err)
	}

	return book, nil

}

func UpdateBook(id string, book Book) (Book, error) {
	db := database.GetDB()

	bookId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	bookUpdated := Book{Id: bookId}

	err = db.Model(&bookUpdated).Where("id = ?", bookId).Updates(Book{BookName: book.BookName, Author: book.Author}).Error

	if err != nil {
		fmt.Println("Error updateing book data :", err)
		return bookUpdated, err
	}

	return bookUpdated, nil

}

// func DeleteBook(id string) int64 {
// 	db := database.GetDB()
// 	defer db.Close()

// 	bookId, err := strconv.Atoi(id)

// 	if err != nil {
// 		panic(err)
// 	}

// 	sqlStatement := `
// 			DELETE FROM books
// 			WHERE id = $1;
// 		`

// 	res, err := db.Exec(sqlStatement, bookId)

// 	if err != nil {
// 		panic(err)
// 	}

// 	count, err := res.RowsAffected()

// 	if err != nil {
// 		panic(err)
// 	}

// 	return count

// }
