package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("failed to connect to db")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {

	createUsersTable := `
		Create table if not exists users(
			id integer primary key autoincrement,
			name text not null,
			email text not null unique,
			password text not null,
			created_at Datetime not null
		)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
		Create table if not exists events(
			id integer primary key autoincrement,
			name text not null,
			description text not null,
			location text not null,
			user_id integer,
			created_at Datetime not null,
			foreign key(user_id) references users(id)
		)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createEventsRegisteratoinsTable := `
		Create table if not exists events_registerations(
			id integer primary key autoincrement,
			user_id integer,
			event_id integer,
			foreign key(user_id) references users(id)
			foreign key(event_id) references events(id)
		)
	`

	_, err = DB.Exec(createEventsRegisteratoinsTable)
	if err != nil {
		panic(err)
	}

}
