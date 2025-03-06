package models

import (
	"time"

	"example/rest/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// func (e Event) Save() error {
// 	fmt.Println("Saving10...")
// 	query := `
// 	INSERT INTO events (name, description, location, dateTime, user_id)
// 	VALUES (?, ?, ?, ?, ?)`

// 	fmt.Println("Saving11...")
// 	stmt, err := db.DB.Prepare(query) //error

// 	fmt.Println("Saving12...")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	fmt.Println("Saving13...")
// 	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Saving14...")
// 	id, err := result.LastInsertId()
// 	fmt.Println("Created ID: ", id)
// 	e.ID = id
// 	return err
// 	return nil
// }

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event = []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil

}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil

}
