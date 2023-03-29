package models

import (
	"database/sql"
	"fmt"
	"strconv"

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
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// memeriksa/verifikasi info psqlInfo benar

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// connecting to database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// var book = Book{}

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
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// memeriksa/verifikasi info psqlInfo benar

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// connecting to database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

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
	// contoh output : [{1 febrianto febri@gmail.com 23 Developer} {2 febrianto febri2@gmail.com 23 Developer} ]
}

func GetBook(id string) Book {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	bookId, err := strconv.Atoi(id)
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
