package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DeleteGroup(id int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM schedule.groups WHERE id=?",
		id)
	if err != nil {
		return err
	}
	return nil
}

type Group struct {
	Id     string
	Degree string
}

func SelectGroups(db *sql.DB) ([]map[string]string, error) {
	result, err := db.Query("SELECT * FROM schedule.groups")
	//defer result.Close()
	if err != nil {
		return nil, err
	}
	var res []map[string]string
	for result.Next() {

		var group Group
		err := result.Scan(&group.Id, &group.Degree)

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, map[string]string{"id": group.Id, "degree": group.Degree})
	}
	return res, nil
}
