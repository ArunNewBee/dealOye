package services

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

var auth smtp.Auth

//MailRequest struct
type MailRequest struct {
	from    string
	to      []string
	subject string
	body    string
}

func (r *MailRequest) SendMail(username, hash, url string) {

	auth = smtp.PlainAuth("", "prabeshnair91@gmail.com", "prajesh1984", "smtp.gmail.com")
	templateData := struct {
		Name string
		URL  string
		HASH string
	}{
		Name: "Prabesh",
		URL:  "http://localhost:9000/emailverification?hash=" + hash + "&username=" + username,
		HASH: hash,
	}
	//r := NewRequest([]string{"junk@junk.com"}, "Hello Junk!", "Hello, World!")
	err := r.ParseTemplate("template.tpl", templateData)

	ok, err := r.SendEmail()
	fmt.Println(ok)
	fmt.Println(err)

}

func NewRequest(to []string, subject, body string) *MailRequest {
	return &MailRequest{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *MailRequest) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "dhanush@geektrust.in", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

//ParseTemplate  parses the html template
func (r *MailRequest) ParseTemplate(templateFileName string, data interface{}) error {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	t, err := template.ParseFiles(pwd + "/dealOye/app/views/EmailTemplates/template.tpl")
	if err != nil {
		fmt.Println(err)
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return err
	}
	r.body = buf.String()
	return nil
}
