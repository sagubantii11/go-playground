package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Venue       string    `json:"venue" binding:"required"`
	Organizer   string    `json:"created_by" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
}

var events = []Event{}

func GetAllEvents() []Event {
	return events
}

func AddEvent(event Event) {
	events = append(events, event)
}