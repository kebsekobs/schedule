package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	api "github.com/kebsekobs/schedule/backend/internal/apientity"
	"github.com/kebsekobs/schedule/backend/internal/generation"
)

// CRUD методы для таблицы "rooms"
func CreateRoom(db *sql.DB, room api.Classroom) error {
	query := "INSERT INTO rooms (id, capacity) VALUES (?, ?)"
	query += " ON DUPLICATE KEY UPDATE capacity = VALUES(capacity)"
	_, err := db.Exec(query, room.ClassroomID, room.Capacity)
	if err != nil {
		return err
	}
	return nil
}

func GetRooms(db *sql.DB) ([]api.Classroom, error) {
	query := "SELECT id, capacity FROM rooms"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []api.Classroom
	for rows.Next() {
		var room api.Classroom
		err := rows.Scan(&room.ClassroomID, &room.Capacity)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func UpdateRoom(db *sql.DB, room api.Classroom) error {
	query := "UPDATE rooms SET capacity = ? WHERE id = ?"
	_, err := db.Exec(query, room.Capacity, room.ClassroomID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRoom(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM rooms WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func InsertRooms(db *sql.DB, rooms []*generation.Room) error {
	query := "INSERT INTO rooms (id, capacity) VALUES "
	for range rooms {
		query += "(?, ?),"
	}
	query = query[:len(query)-1] // Remove the trailing comma
	query += " ON DUPLICATE KEY UPDATE capacity = VALUES(capacity)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var values []interface{}
	for _, room := range rooms {
		values = append(values, room.ID, room.Capacity)
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}

	return nil
}

func SelectRooms(db *sql.DB) ([]*generation.Room, error) {
	var rooms []*generation.Room

	rows, err := db.Query("SELECT id, capacity FROM rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var capacity int
		if err := rows.Scan(&id, &capacity); err != nil {
			return nil, err
		}

		rooms = append(rooms, &generation.Room{ID: id, Capacity: capacity})
	}

	return rooms, nil
}
