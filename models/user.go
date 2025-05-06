package models

import (
	"log"
	"todolist/db"
)

type User struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Id       int    `json:"id"`
}

func AddUser() {
	var user User
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Email, user.Password)
}
