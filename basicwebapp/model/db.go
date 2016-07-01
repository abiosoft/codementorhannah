package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	log.Fatal(err)
}
