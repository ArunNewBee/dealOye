package models

import (
	"dealOye/app"
	"fmt"
	"time"

	//"golang.org/x/crypto/bcrypt"
)

type Register struct {
	Lname    string
	Fname    string
	Email    string
	PNumber  string
	Pass     string
	Type     string
	UserName string
}

//DoRegistration function to insert user data
func (r Register) DoRegistration() {
	date := time.Now().Format("2006-01-02 15:04:05")
	// password := []byte(r.Pass)
	// // Hashing the password with the default cost of 10
	// hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	sql := "INSERT INTO `users` (lName,fName,userName,email,password,createdOn,lastLogin,phone) VALUES (?,?,?,?,?,?,?,?)"
	_, err := app.DB.Exec(sql, r.Lname, r.Fname, r.UserName, r.Email, r.Pass, date, date, r.PNumber)
	if err != nil {
		fmt.Println(err)
	}
}

func CheckUserName(user string) bool {
	var count int
	sql1 := "SELECT COUNT(*) FROM `users` WHERE username=?"
	err := app.DB.QueryRow(sql1, user).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	if count > 0 {
		return false
	}
	return true
}
func CheckEmail(email string) bool {
	var count int
	sql := "SELECT COUNT(*) FROM `users` WHERE email=?"
	err := app.DB.QueryRow(sql, email).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	if count > 0 {
		return false
	}
	return true
}
