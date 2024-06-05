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

/*
// CRUD методы для таблицы "groups"
func createGroup(id string, name string, quantity int) {
    query := "INSERT INTO groups (id, name, quantity) VALUES (?, ?, ?)"
    _, err := db.Exec(query, id, name, quantity)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Group created successfully")
}

func getGroups() []Group {
    query := "SELECT id, name, quantity FROM groups"
    rows, err := db.Query(query)
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    var groups []Group
    for rows.Next() {
        var group Group
        err := rows.Scan(&group.ID, &group.Name, &group.Quantity)
        if err != nil {
            panic(err.Error())
        }
        groups = append(groups, group)
    }
    return groups
}

func updateGroupName(id string, newName string) {
    query := "UPDATE groups SET name = ? WHERE id = ?"
    _, err := db.Exec(query, newName, id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Group name updated successfully")
}

func deleteGroup(id string) {
    query := "DELETE FROM groups WHERE id = ?"
    _, err := db.Exec(query, id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Group deleted successfully")
}

// CRUD методы для таблицы "rooms"
func createRoom(id string, capacity int) {
    query := "INSERT INTO rooms (id, capacity) VALUES (?, ?)"
    _, err := db.Exec(query, id, capacity)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Room created successfully")
}

func getRooms() []Room {
    query := "SELECT id, capacity FROM rooms"
    rows, err := db.Query(query)
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    var rooms []Room
    for rows.Next() {
        var room Room
        err := rows.Scan(&room.ID, &room.Capacity)
        if err != nil {
            panic(err.Error())
        }
        rooms = append(rooms, room)
    }
    return rooms
}

func updateRoomCapacity(id string, newCapacity int) {
    query := "UPDATE rooms SET capacity = ? WHERE id = ?"
    _, err := db.Exec(query, newCapacity, id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Room capacity updated successfully")
}

func deleteRoom(id string) {
    query := "DELETE FROM rooms WHERE id = ?"
    _, err := db.Exec(query, id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Room deleted successfully")
}

// CRUD методы для таблицы "teachers"
func createTeacher(name string) {
    query := "INSERT INTO teachers (name) VALUES (?)"
    _, err := db.Exec(query, name)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Teacher created successfully")
}

func getTeachers() []Teacher {
    query := "SELECT id, name FROM teachers"
    rows, err := db.Query(query)
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    var teachers []Teacher
    for rows.Next() {
        var teacher Teacher
        err := rows.Scan(&teacher.ID, &teacher.Name)
        if err != nil {
            panic(err.Error())
        }
        teachers = append(teachers, teacher)
    }
    return teachers
}

// Продолжайте реализовывать методы update и delete для таблицы "teachers" по аналогии с вышеприведенными примерами
// CRUD методы для таблицы "teachers"
func updateTeacherName(id string, newName string) {
    query := "UPDATE teachers SET name = ? WHERE id = ?"
    _, err := db.Exec(query, newName, id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Teacher name updated successfully")
}

func deleteTeacher(id string) {
    query := "DELETE FROM teachers WHERE id = ?"
    _, err := db.Exec(query, id)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Teacher deleted successfully")
}
*/
