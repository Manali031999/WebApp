package db

import (
	"fmt"
	"database/sql"
)


func DbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/form")
    if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected Successfully")
    return db
}
