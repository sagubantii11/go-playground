package models

import (
	"database/sql"
	"time"

	"github.com/sagubantii11/go-playground/sqldb"
)

type Event struct {
	ID          int64       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Venue       string    `json:"venue" binding:"required"`
	Organizer   string    `json:"created_by" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
}

var events = []Event{}

func GetAllEvents() []Event {
	getQuery := `SELECT * FROM events`
	rows, err := sqldb.ExecuteSelect(getQuery)
	if err != nil {
		panic(err)
	}
	events = rowsToEvents(rows)
	defer rows.Close()
	return events
}

func rowsToEvents(rows *sql.Rows) []Event {
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Venue, &event.Organizer, &event.DateTime)
		if err != nil {
			panic(err)
		}
		events = append(events, event)
	}
	return events
}

func AddEvent(event Event) {
	insertSQLStatement := `
	INSERT INTO events (title, description, venue, created_by, date_time)
	VALUES (?, ?, ?, ?, ?)
	`
	id, err := sqldb.ExecuteInserts(insertSQLStatement, event.Title, event.Description, event.Venue, event.Organizer, event.DateTime)
	if err != nil {
		panic(err)
	}
	event.ID = id
}