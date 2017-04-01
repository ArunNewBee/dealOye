package controllers

import (
	"crypto/rand"
	"dealOye/app/models"
	"dealOye/app/routes"
	"dealOye/app/services"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/revel/revel"
)

type Register struct {
	*revel.Controller
}
type DuplicateEmailValidator struct{}

func (v DuplicateEmailValidator) IsSatisfied(email interface{}) bool {

	if models.CheckEmail(email.(string)) {
		return true
	}
	return false

}

func (v DuplicateEmailValidator) DefaultMessage() string {
	return "Sorry, that email is already in use"
}

type DuplicateUserNameValidator struct{}

func (v DuplicateUserNameValidator) IsSatisfied(userName interface{}) bool {
	if models.CheckUserName(userName.(string)) {
		return true
	}
	return false
}

func (v DuplicateUserNameValidator) DefaultMessage() string {
	return "Sorry, that user name is already in taken"
}

// //username Checker
// type emailChecker struct{}

// func (u emailChecker) IsSatisified(i interface{}) bool {
// 	s, k := i.(string)

// 	if !k {
// 		return false
// 	}

// 	return true
// }

// func (u emailChecker) DefaultMessage() string {
// 	return "username already in use"
// }

func (c Register) GetRegister() revel.Result {
	return c.Render()
}

func (c Register) DoRegister() revel.Result {
	// var userNameCheck usernameChecker
	fname := c.Params.Get("fname")
	lname := c.Params.Get("lname")
	email := c.Params.Get("email")
	phone := c.Params.Get("phone")
	pass := c.Params.Get("pass")
	cpass := c.Params.Get("cpass")
	username := c.Params.Get("username")
	c.Validation.Required(fname).Message("Please enter a username")
	c.Validation.Required(phone).Message("Mobile number is required")
	c.Validation.Required(pass == cpass).Message("Your passwords do not match.")
	c.Validation.Check(username, revel.Required{}, DuplicateUserNameValidator{})
	c.Validation.Check(email, revel.Required{}, DuplicateEmailValidator{})
	var register models.Register
	register.Fname = fname
	register.Lname = lname
	register.Email = email
	register.PNumber = phone
	register.Pass = pass
	register.UserName = username
	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Register.GetRegister)
	}
	register.EmailExpiry = time.Now().Unix() + 60*60

	hash, err := GenerateRandomString(32)
	if err != nil {
		fmt.Println(err)
	}
	register.Emailhash = hash

	//register.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(register.Pass), bcrypt.DefaultCost)
	err = register.DoRegistration()
	if err != nil {
		fmt.Println(err)
	}

	c.Session["user"] = register.UserName
	c.Flash.Success("Welcome, " + register.Fname)

	mail := services.NewRequest([]string{register.Email}, "DealOyeDeal Mail Confirmation", "")
	mail.SendMail(register.UserName, register.Emailhash, "ramseyRmas")
	return c.Redirect(routes.Application.Index())
}

func (c Register) EmailVerification(username, hash string) revel.Result {

	msg, err := models.EmailVerify(username, hash)

	if err != nil {
		fmt.Println(err)
	}
	if msg == "emailExpired" {
		fmt.Println(msg)
	}
	return c.Render(msg)
}
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
