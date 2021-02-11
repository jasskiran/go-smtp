package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	send("hello there")
}

func send(body string) {
	from := os.Getenv("EMAIL_FROM")
	pass := os.Getenv("EMAIL_PASSWORD")
	to := os.Getenv("EMAIL_TO")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail(os.Getenv("ADDRESS"),
		smtp.PlainAuth("", from, pass, os.Getenv("HOST")),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("sent", msg)
}

