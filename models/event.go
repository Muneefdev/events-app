package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e *Event) Save() {
	events = append(events, *e)
}

func (e *Event) Update() {
}

func (e *Event) Delete() {
}

func New() *Event {
	return &Event{}
}

func GetEvent() *Event {
	return nil
}

func GetAllEvents() *[]Event {
	return nil
}
