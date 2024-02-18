package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteRoom(id int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM rooms WHERE id=?",
		id)
	if err != nil {
		return err
	}
	return nil
}
