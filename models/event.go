package models

import (
	"time"

	"example.com/learngo/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"dateTime"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, user_id, created_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		println("Error preparing statement")
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.CreatedAt)

	if err != nil {
		println("Error executing statement")
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		println("Error getting last insert id")
		return err
	}

	e.ID = id

	return nil

}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID, &event.CreatedAt)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}
