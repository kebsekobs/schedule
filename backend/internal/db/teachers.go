package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kebsekobs/schedule/backend/internal/generation"
)

// func InsertTeachers(db *sql.DB, teachers map[string]*generation.Teacher) error {
// 	query := "INSERT IGNORE INTO teachers (name) VALUES "
// 	for range teachers {
// 		query += "(?),"
// 	}
// 	query = query[:len(query)-1] // Remove the trailing comma

// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	var values []interface{}
// 	for _, teacher := range teachers {
// 		values = append(values, teacher.Name)
// 	}

// 	_, err = stmt.Exec(values...)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func InsertTeachers(db *sql.DB, teachers map[string]*generation.Teacher) (map[string]*generation.Teacher, error) {
	query := "INSERT IGNORE INTO teachers (name) VALUES (?)"

	for _, teacher := range teachers {
		stmt, err := db.Prepare(query)
		if err != nil {
			return nil, err
		}

		res, err := stmt.Exec(teacher.Name)
		if err != nil {
			return nil, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}

		teacher.ID = int(id)

		defer stmt.Close()
	}

	return teachers, nil
}

func SelectTeachers(db *sql.DB) ([]*generation.Teacher, error) {
	var teachers []*generation.Teacher

	rows, err := db.Query("SELECT id, name FROM teachers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		teachers = append(teachers, &generation.Teacher{ID: id, Name: name})
	}

	return teachers, nil
}
