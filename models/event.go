package models

import (
	"database/sql"
	"time"

	"github.com/sagubantii11/go-playground/sqldb"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Venue       string    `json:"venue" binding:"required"`
	Organizer   string    `json:"created_by" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
}

var events = []Event{}

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


func GetEventByID(id int64) Event {
	getQuery := `SELECT * FROM events WHERE id = ?;`
	rows, err := sqldb.ExecuteSelect(getQuery, id)
	if err != nil {
		panic(err)
	}
	events = rowsToEvents(rows)
	defer rows.Close()
	if len(events) == 0 {
		return Event{}
	}
	return events[0]
}

func AddEvent(event Event) Event {
	insertSQLStatement := `
	INSERT INTO events (title, description, venue, created_by, date_time)
	VALUES (?, ?, ?, ?, ?);
	`
	id, err := sqldb.ExecuteInserts(insertSQLStatement, event.Title, event.Description, event.Venue, event.Organizer, event.DateTime)
	if err != nil {
		panic(err)
	}
	event.ID = id
	return event
}

func UpdateEventByID(id int64, event Event) Event {
	updateSQLStatement := `
	UPDATE events
	SET title = ?, description = ?, venue = ?, created_by = ?, date_time = ?
	WHERE id = ?;
	`
	_, err := sqldb.ExecuteInserts(updateSQLStatement, event.Title, event.Description, event.Venue, event.Organizer, event.DateTime, id)
	if err != nil {
		panic(err)
	}
	// event.ID, _ = event.ID, err
	return event
}

func DeleteEventByID(id int64) {
	deleteSQLStatement := `DELETE FROM events WHERE id = ?;`
	_, err := sqldb.ExecuteInserts(deleteSQLStatement, id)
	if err != nil {
		panic(err)
	}
}