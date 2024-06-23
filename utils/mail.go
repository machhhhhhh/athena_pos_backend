package utils

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
	// gomail "gopkg.in/mail.v2"
)

type MailPayload struct {
	To      string
	Subject string
	Body    string
}

type RefEmail struct {
	Subject string
	Name    string
	RefName string
	Money   string
	Creater string
	Detail  string
	Link    string
	Section string
}

func SendEmail(payload *MailPayload) error {

	// if payload.To != "" && IsValidEmail(payload.To) {

	// 	mail_host := os.Getenv("E_MEMO_MAIL_HOST")
	// 	mail_port := os.Getenv("E_MEMO_MAIL_PORT")
	// 	mail_username := os.Getenv("E_MEMO_MAIL_USERNAME")
	// 	mail_password := os.Getenv("E_MEMO_MAIL_PASSWORD")
	// 	mail_from := os.Getenv("E_MEMO_MAIL_FROM")

	// 	if mail_host == "" || mail_port == "" || mail_from == "" {
	// 		return errors.New("environment is not set")
	// 	}

	// 	port, err := strconv.Atoi(mail_port)
	// 	if err != nil {
	// 		return errors.New("port ของ mail เป็นได้แค่ตัวเลขเท่านั้น")
	// 	}

	// 	mail := gomail.NewMessage()

	// 	// Set E-Mail sender
	// 	mail.SetHeader("From", mail_from)

	// 	// Set E-Mail receivers
	// 	mail.SetHeader("To", payload.To)

	// 	// Set E-Mail subject
	// 	mail.SetHeader("Subject", payload.Subject)

	// 	// Set E-Mail body. You can set plain text or html with text/html
	// 	mail.SetBody("text/html", payload.Body)

	// 	// Settings for SMTP server
	// 	diabler := gomail.NewDialer(mail_host, port, mail_username, mail_password)
	// 	// diabler := gomail.NewDialer(mail_host, port, "", "")

	// 	// This is only needed when SSL/TLS certificate is not valid on server.
	// 	// In production this should be set to false.
	// 	diabler.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 	if err := diabler.DialAndSend(mail); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func SendEmailWithTemplate(ref RefEmail, mail *MailPayload, tmpl *template.Template) {
	buff := new(bytes.Buffer)
	tmpl.Execute(buff, ref)

	mail.Body = buff.String()

	err := SendEmail(mail)
	if err != nil {
		panic(err)
	}
}

func GetEmailTemplate(file_name string) *template.Template {
	// find root
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// get file path
	root_path := fmt.Sprintf(strings.Replace(root, `\`, "/", -1) + "/templates/" + file_name + ".html")

	//parse file from directory.
	tmpl, err := template.ParseFiles(root_path)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func IsValidEmail(email string) bool {
	// Define a regular expression pattern for a basic email validation
	// This regex pattern is a simplified version and may not cover all edge cases.
	// For a more comprehensive email validation, you can use a more complex regex pattern.
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)

	// Use the regex to match the email string
	return regex.MatchString(email)
}
