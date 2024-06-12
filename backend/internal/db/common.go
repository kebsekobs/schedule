package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func truncateTable(db *sql.DB, tableName string) error {
	_, err := db.Exec("TRUNCATE TABLE " + tableName)
	if err != nil {
		return err
	}
	return nil
}
