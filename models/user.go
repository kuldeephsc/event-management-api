package models

import (
	"errors"

	"github.com/kuldeephsc/api/db"
	"github.com/kuldeephsc/api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordValid {
		return errors.New("Invalid credentials")
	}

	return nil

}

func (e *Event) RegisterUserForEvent(userId int64) error {
	query := `INSERT INTO event_registrations(event_id, user_id)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err1 := stmt.Exec(e.ID, userId)
	return err1
}

func (e *Event) UnregisterUserFromEvent(userId int64) error {
	query := `DELETE FROM event_registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err1 := stmt.Exec(e.ID, userId)
	return err1
}
