package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitiateDatabaseConnection() *DB {
	db, err := sql.Open("mysql", "root:Astro@123@tcp(127.0.0.1:3306)/goBackend")

	if err != nil {
		panic(err.Error())
	}

	return db
}
