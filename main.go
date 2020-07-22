package main

import (
	"log"
	"net/smtp"
)

func main() {
	addr := "smtp.gmail.com:587"
	host := "smtp.gmail.com"
	identity := ""
	from := "" // required
	password := "" // required
	to := "" // required
	subject := "This is an example email"
	body := "Hello"
	msg := "From:" + from + "\r\n" + "To:" + to + "\r\n" + "Subject:" + subject + "\r\n" + body

	err := smtp.SendMail(
		addr,
		smtp.PlainAuth(identity, from, password, host),
		from,
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		log.Println(err)
	}
}
