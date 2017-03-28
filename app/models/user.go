package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId             int
	Name               string
	Username, Password string
	HashedPassword     []byte
}

func GetUser(username string) *User {
	var user User
	user.Username = "prabesh"
	pa := "prabesh"
	pass := []byte(pa)
	user.HashedPassword, _ = bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	return &user
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}
