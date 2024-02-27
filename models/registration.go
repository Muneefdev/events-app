package models

import "github.com/muneefdev/events-app/db"

type Registration struct {
	ID      int   `json:"id"`
	UserID  int64 `json:"user_id"`
	EventID int64 `json:"event_id"`
}

func NewRegistration(userID, eventID int64) *Registration {
	return &Registration{
		UserID:  userID,
		EventID: eventID,
	}
}

func (r *Registration) RegisterEvent() error {
	query := `
  INSERT INTO registrations (user_id, event_id) 
  VALUES (?, ?)
  `
	_, err := db.DB.Exec(query, r.UserID, r.EventID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEventRegistration(registrationID int64) error {
	query := `
  DELETE FROM registrations
  WHERE id = ?
  `
	_, err := db.DB.Exec(query, registrationID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllRegistrations() ([]Registration, error) {
	query := `
  SELECT * FROM registrations
  `
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var registrations []Registration
	for rows.Next() {
		var registration Registration
		err = rows.Scan(&registration.ID, &registration.UserID, &registration.EventID)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}

	return registrations, nil
}

func GetRegistrationByID(registrationID int64) (*Registration, error) {
	query := `
  SELECT * FROM registrations
  WHERE id = ?
  `
	row := db.DB.QueryRow(query, registrationID)
	var registration Registration
	err := row.Scan(&registration.ID, &registration.UserID, &registration.EventID)
	if err != nil {
		return nil, err
	}
	return &registration, nil
}

func CheckRegistration(userID, eventID int64) bool {
	query := `
  SELECT * FROM registrations
  WHERE user_id = ? AND event_id = ?
  `
	row := db.DB.QueryRow(query, userID, eventID)
	var registration Registration
	err := row.Scan(&registration.ID, &registration.UserID, &registration.EventID)
	if err != nil {
		return false
	}

	return true
}
