package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

var smtpPassword = os.Getenv("SMTP_PASSWORD")
var smtpFrom = os.Getenv("SMTP_FROM")

var bodyTemplate = `
  Greetings Elf 👋

  Your secret santa match is %s.

  Merry Christmas! 🎅

  - Secrit Santa
`

// SendSecritSantaMatch sends the match result for a user.
func SendSecritSantaMatch(match Match) {
	msg := "From: Secrit Santa" + "\n" +
		"To: " + match.Santa + "\n" +
		"Subject: Merry Christmas Son\n\n" +
		fmt.Sprintf(bodyTemplate, match.Receiver)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", smtpFrom, smtpPassword, "smtp.gmail.com"),
		smtpFrom,
		[]string{match.Santa},
		[]byte(msg),
	)

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Println("Sent match to ", match.Santa)
}
