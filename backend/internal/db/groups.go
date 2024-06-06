package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	api "github.com/kebsekobs/schedule/backend/internal/apientity"
	"github.com/kebsekobs/schedule/backend/internal/generation"
)

// CRUD методы для таблицы "groups"
func CreateGroup(db *sql.DB, group api.Group) error {
	query := "INSERT INTO groups (id, name, quantity) VALUES (?, ?, ?)"
	query += " ON DUPLICATE KEY UPDATE name, quantity = VALUES(name, quantity)"
	_, err := db.Exec(query, group.ID, group.Name, group.Capacity)
	if err != nil {
		return err
	}
	return nil
}

func GetGroups(db *sql.DB) ([]api.Group, error) {
	query := "SELECT id, name, quantity FROM groups"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []api.Group
	for rows.Next() {
		var group api.Group
		err := rows.Scan(&group.ID, &group.Name, &group.Capacity)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func UpdateGroup(db *sql.DB, group api.Group) error {
	query := "UPDATE groups SET name = ?, quantity = ? WHERE id = ?"
	_, err := db.Exec(query, group.Name, group.Capacity, group.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGroup(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM schedule.groups WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

// type Group struct {
// 	Id     string
// 	Degree string
// }

// func SelectGroups(db *sql.DB) ([]map[string]string, error) {
// 	result, err := db.Query("SELECT * FROM schedule.groups")
// 	//defer result.Close()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var res []map[string]string
// 	for result.Next() {

// 		var group Group
// 		err := result.Scan(&group.Id, &group.Degree)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		res = append(res, map[string]string{"id": group.Id, "degree": group.Degree})
// 	}
// 	return res, nil
// }

func InsertGroups(db *sql.DB, groups map[string]*generation.Group) error {
	query := "INSERT INTO groups (id, quantity) VALUES "
	for range groups {
		query += "(?, ?),"
	}
	query = query[:len(query)-1] // Remove the trailing comma
	query += " ON DUPLICATE KEY UPDATE quantity = VALUES(quantity)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var values []interface{}
	for _, group := range groups {
		values = append(values, group.ID, group.Quantity)
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}

	return nil
}

func SelectGroups(db *sql.DB) ([]*generation.Group, error) {
	var groups []*generation.Group

	rows, err := db.Query("SELECT id, quantity FROM groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var quantity int
		var id string
		if err := rows.Scan(&id, &quantity); err != nil {
			return nil, err
		}

		groups = append(groups, &generation.Group{ID: id, Quantity: quantity})
	}

	return groups, nil
}
