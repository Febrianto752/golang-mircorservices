package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin123"
	dbname   = "golang-dasar"
)

// var (
// 	db  *sql.DB
// 	err error
// )

func GetDB() *sql.DB {
	var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("db :", db)

	fmt.Println("Successfully connected to database")

	return db
}
