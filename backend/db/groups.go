package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteGroup(id int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM groups WHERE id=?",
		id)
	if err != nil {
		return err
	}
	return nil
}
