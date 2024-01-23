package models

import (
	"errors"

	"vs45tech.com/event/db"
	"vs45tech.com/event/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() (int64, error) {
	query := `INSERT INTO users(email,password)
	VALUES(?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	hashPassWord, err := utils.HashPassWord(u.Password)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(u.Email, hashPassWord)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	u.ID = id
	//events = append(events, e)
	return id, err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)
	if err != nil {
		return errors.New("Invalid Credentials")
	}
	passwordIsValid := utils.CheckPasswordHashed(u.Password, retrievePassword)
	if !passwordIsValid {
		return errors.New("Invalid Credentials")
	}
	return nil
}
