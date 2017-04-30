package models

import (
	"dealOye/app"
	"fmt"
	"time"
)

//Register users
type Register struct {
	Lname          string
	Fname          string
	Email          string
	PNumber        string
	Pass           string
	Type           string
	UserName       string
	HashedPassword []byte
	EmailExpiry    int64
	Emailhash      string
}

//DoRegistration function to insert user data
func (r Register) DoRegistration() error {
	date := time.Now().Format("2006-01-02 15:04:05")
	// password := []byte(r.Pass)
	// Hashing the password with the default cost of 10
	// hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	// if err != nil {
	// 	return err
	// }
	sql := "INSERT INTO `users` (lName,fName,userName,email,password,createdOn,lastLogin,phone,emailhash,emailExpiry) VALUES (?,?,?,?,?,?,?,?,?,?)"
	_, err := app.DB.Exec(sql, r.Lname, r.Fname, r.UserName, r.Email, r.Pass, date, date, r.PNumber, r.Emailhash, r.EmailExpiry)
	if err != nil {
		return err
	}

	// _, err := app.DB.Exec("DELETE FROM `users` WHERE username=? OR email=?", "prabesh2321", "prabeshnair91@gmail.com")
	// if err != nil {
	// 	return err
	// }
	return nil
}

//CheckUserName checks if username exists in db or not
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

//CheckEmail checks if email exists in db or not
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

func EmailVerify(username, hash string) (string, error) {
	var count, id int64
	sql := "SELECT emailExpiry,id FROM `users` WHERE username=? and emailhash=?"
	err := app.DB.QueryRow(sql, username, hash).Scan(&count, &id)
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now().Unix()
	if now-count > 60*60 {
		return "emailExpired", nil
	}

	upd := "UPDATE `users` SET emailVerified=? WHERE id=?"
	_, err = app.DB.Exec(upd, "approved", id)
	if err != nil {
		fmt.Println(err)
	}
	return "success", nil
}
