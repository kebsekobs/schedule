package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	api "github.com/kebsekobs/schedule/backend/internal/apientity"
	"github.com/kebsekobs/schedule/backend/internal/generation"
)

// CRUD методы для таблицы "teachers"
func CreateTeacher(db *sql.DB, teacher api.Teacher) error {
	query := "INSERT IGNORE INTO teachers (name) VALUES (?)"
	_, err := db.Exec(query, teacher.Name)
	if err != nil {
		return err
	}
	return nil
}

func GetTeachers(db *sql.DB) ([]api.Teacher, error) {
	query := "SELECT id, name FROM teachers"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []api.Teacher
	for rows.Next() {
		var teacher api.Teacher
		err := rows.Scan(&teacher.ID, &teacher.Name)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func GetTeacherByID(db *sql.DB, id int) (*api.Teacher, error) {
	query := "SELECT id, name FROM teachers WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teacher api.Teacher
	for rows.Next() {
		var teacher api.Teacher
		err := rows.Scan(&teacher.ID, &teacher.Name)
		if err != nil {
			return nil, err
		}
	}
	return &teacher, nil
}

func UpdateTeacher(db *sql.DB, teacher api.Teacher) error {
	query := "UPDATE teachers SET name = ? WHERE id = ?"
	_, err := db.Exec(query, teacher.Name, teacher.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTeacher(db *sql.DB, id int) error {
	query := "DELETE FROM teachers WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func InsertTeachers(db *sql.DB, teachers map[string]*generation.Teacher) (map[string]*generation.Teacher, error) {
	err := truncateTable(db, "teachers")
	if err != nil {
		return nil, err
	}
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

/*

 */
