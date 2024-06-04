package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kebsekobs/schedule/backend/internal/generation"
)

func DeleteRoom(id int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM rooms WHERE id=?",
		id)
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
