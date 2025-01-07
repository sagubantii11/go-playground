package models

import (
	"github.com/sagubantii11/go-playground/sqldb"
	"github.com/sagubantii11/go-playground/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	EmailID  string `json:"email_id" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
}

func (u User) RegisterUser() {
	query := `INSERT INTO users (username, password, email_id, full_name) VALUES (?, ?, ?, ?);`
	hPass, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
	}

	_, err = sqldb.DB.Exec(query, u.Username, hPass, u.EmailID, u.FullName)
	if err != nil {
		panic(err)
	}
}

func (u User) VerifyUserLogin() bool {
	query := `SELECT * FROM users WHERE username = ?;`
	row := sqldb.DB.QueryRow(query, u.Username)
	var dbUser User
	err := row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password, &dbUser.EmailID, &dbUser.FullName)
	if err != nil {
		panic(err)
	}

	err = utils.CheckPasswordHash(u.Password, dbUser.Password)
	
	return err == nil
}