package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func connectDatabse() {
	//db, err = sql.Open("mysql", "root:password@/test")
	//db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	//the error only returns error when sqldriver name is wrong, otherwise not.
	db, err = sql.Open("mysql", "root:C&C@2968!@tcp(localhost)/clickncare")
	if err == nil {
		fmt.Println("Database connected.")
	}

}
