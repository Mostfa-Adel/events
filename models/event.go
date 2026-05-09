package models

import (
	"errors"
	"time"

	"github.com/go-events/db"
)

type scannable interface {
	Scan(...any) error
}

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	CreatedAt   time.Time
	UserID      int64
}

func NewEvent(name, description, location string, userId int64) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		UserID:      userId,
	}
}

func GetAllEvents() ([]Event, error) {
	query := "select * from events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Event
	for rows.Next() {
		row, err := scanEvent(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, *row)
	}
	return results, nil
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events 
		(name, description, location, user_id, created_at)
		values
		(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	e.CreatedAt = time.Now()
	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserID, e.CreatedAt)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()

	if err != nil {
		return err
	}
	e.ID = id
	return err
}

func GetEvent(id int64) (*Event, error) {
	query := "select * from events where id=?"
	row := db.DB.QueryRow(query, id)
	event, err := scanEvent(row)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (event *Event) DeleteEvent() error {
	query := "delete from events where id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}

func scanEvent(row scannable) (*Event, error) {
	var event Event
	return &event, row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.UserID,
		&event.CreatedAt,
	)
}

func (event *Event) UpdateEvent() error {
	query := `update events set name=?, description=?, location=? where id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.ID)
	return err
}

func (event *Event) RegisterUserInEvent(userId int64) error {
	query := `insert into events_registerations ( user_id, event_id) values (?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("couldnt register ")
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, event.ID)
	return err
}

func (event *Event) DeleteUser(userId int64) error {
	query := `delete from events_registerations where event_id=? and user_id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("failed to cancel")
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID, userId)
	return err
}
