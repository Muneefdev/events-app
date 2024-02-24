package models

import (
	"time"

	"example.com/rest-apis/events/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events (
  name, description, location, dateTime, user_id) 
  VALUES (?, ?, ?, ?, ?) 
  `
	sqlStmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer sqlStmt.Close()

	result, err := sqlStmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.UserId = id
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

