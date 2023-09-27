package main

import (
	"database/sql"
	"fmt"
	"github.com/kebsekobs/schedule/tree/main/backend/config"
	"github.com/kebsekobs/schedule/tree/main/backend/db"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	_db = dbConnect()
)

// конектимся к бд
func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", config.DBConfig.GetString("user")+":"+config.DBConfig.GetString(
		"password")+"@/"+config.DBConfig.GetString("name"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected successfully.")
	}
	// настройка пула соединений
	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	return db
}
func main() {
	a, b := db.SelectGroups(_db)
	fmt.Println(a, b)
}
