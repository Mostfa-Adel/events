package models

import (
	"errors"
	"time"

	"github.com/go-events/db"
	"github.com/go-events/utils"
)

type User struct {
	ID        int64
	Name      string
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	CreatedAt time.Time
}

func (user *User) SaveUser() error {
	query := "insert into users(name, email, password, created_at) values (?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.Hash(user.Password)
	if err != nil || hashedPassword == "" {
		return errors.New("failed password")
	}
	res, err := stmt.Exec(user.Name, user.Email, hashedPassword, time.Now())
	if err != nil {
		return err
	}
	user.ID, err = res.LastInsertId()
	return err
}

func ValidateUserCredintials(email, password string) (int64, error) {
	row := db.DB.QueryRow("select id,password from users where email=?", email)

	var hashedPassword string
	var userId int64
	err := row.Scan(&userId, &hashedPassword)
	if err != nil {
		return 0, errors.New("credintials not matching")
	}
	isMatch := utils.IsPasswordMatch(hashedPassword, password)
	if isMatch {
		return 0, errors.New("credintials not matching")
	}
	return userId, nil
}

// func GetUserByEmail(email string) (user *User, error) {
// 	query := "select 1 from users where email=?"
// 	row := db.DB.QueryRow(query, email)
// 	row.Scan()
// }
