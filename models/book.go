package models

import (
	"database/sql"
	"fmt"
	"strconv"

	"golang_microservices/database"

	_ "github.com/lib/pq"
)

type Book struct {
	Id          int    `json:"book_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin123"
	dbname   = "golang-dasar"
)

var (
	db  *sql.DB
	err error
)

func CreateBook(book Book) {
	db := database.GetDB()
	fmt.Println("db :", db)
	defer db.Close()

	sqlStatement := `
    INSERT INTO books (title, author, description) VALUES ($1, $2, $3)
    returning *
  `

	err = db.QueryRow(sqlStatement, book.Title, book.Author, book.Description).Scan(&book.Id, &book.Title, &book.Author, &book.Description)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Book Data : %+v \n", book)
}

func GetBooks() []Book {
	db := database.GetDB()

	var results = []Book{}

	sqlStatement := `SELECT * FROM books`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		fmt.Println("panic 1")
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Description)

		if err != nil {
			fmt.Println("Panic 2")
			panic(err)
		}

		results = append(results, book)
	}

	return results
}

func GetBook(id string) Book {
	db := database.GetDB()

	bookId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	var book Book

	sqlStatement := `SELECT * FROM books WHERE id = $1`

	rows, err := db.Query(sqlStatement, bookId)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Description)
		if err != nil {
			panic(err)
		}

	}

	return book

}

func UpdateBook(id string, book Book) int64 {
	db := database.GetDB()

	bookId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	sqlStatement := `
			UPDATE books 
			SET title = $2, author = $3, description = $4
			WHERE id = $1;
		`

	res, err := db.Exec(sqlStatement, bookId, book.Title, book.Author, book.Description)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	return count

}

func DeleteBook(id string) int64 {
	db := database.GetDB()

	bookId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	sqlStatement := `
			DELETE FROM books
			WHERE id = $1;
		`

	res, err := db.Exec(sqlStatement, bookId)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	return count

}
