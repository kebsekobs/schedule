package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kebsekobs/schedule/backend/iternal/config"
	"github.com/kebsekobs/schedule/backend/iternal/db"

	_ "github.com/go-sql-driver/mysql"
)

var (
	_db = dbConnect()
)

// конектимся к бд
func dbConnect() *sql.DB {
	fmt.Println(config.DBConfig.GetString("user"))
	localdb, err := sql.Open("mysql", config.DBConfig.GetString("user")+":"+config.DBConfig.GetString(
		"password")+"@/"+config.DBConfig.GetString("name"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected successfully.")
	}
	// настройка пула соединений
	localdb.SetConnMaxLifetime(time.Minute * 60)
	localdb.SetMaxOpenConns(1)
	localdb.SetMaxIdleConns(1)
	return localdb
}

func main() {
	a, b := db.SelectGroups(_db)
	fmt.Println(a, b)
	//db.DeleteGroup(101, _db)
}
