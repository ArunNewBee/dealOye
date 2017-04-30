package models

import (
	"dealOye/app"
	"fmt"
)

type User struct {
	UserId             int
	Name               string
	Username, Password string
	HashedPassword     []byte
}

func GetUser(username string) *User {
	var user User
	var password string
	sql := "SELECT password FROM `users` WHERE username=?"
	err := app.DB.QueryRow(sql, username).Scan(&password)

	if err != nil {
		fmt.Println(err)
	}

	user.Password = password
	return &user
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}
