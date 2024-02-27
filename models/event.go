package models

import (
	"time"

	"github.com/muneefdev/events-app/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"dateTime"`
	UserId      int64     `json:"userId"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events (
  name, description, location, dateTime, user_id) 
  VALUES (?, ?, ?, ?, ?) 
  `
	result, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	e.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) UpdateEvent() error {
	sqlStmt := `UPDATE events
  SET name = ?, description = ?, location = ?, dateTime = ?
  WHERE id = ?`
	_, err := db.DB.Exec(sqlStmt, e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) Delete() {
}

func New() *Event {
	return &Event{}
}

func GetEventById(id int64) (*Event, error) {
	sqlStmt := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(sqlStmt, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description,
		&event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() (*[]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID, &event.Name, &event.Description, &event.Location,
			&event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return &events, nil
}

func DeleteEventById(id int64) error {
	sqlStmt := `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(sqlStmt, id)
	if err != nil {
		return err
	}
	return nil
}
