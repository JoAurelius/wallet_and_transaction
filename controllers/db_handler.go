package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_quiz1_pbp?parseTime=true&loc=Asia%2fJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
