package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/gomail.v2"
)

type application struct {
	client *gomail.Dialer
}

type Request struct {
	Message string `json:"message"`
}

func main() {
	d := gomail.NewDialer("mail.tysonjenkins.dev", 587, "timmy@tysonjenkins.dev", "password")
	d.TLSConfig = &tls.Config{ServerName: "mail.tysonjenkins.dev", InsecureSkipVerify: true}
	d.LocalName = "mail.tysonjenkins.dev"

	app := application{client: d}

	http.HandleFunc("/", app.Send)

	fmt.Println("Listening on localhost:8080")
	log.Panic(http.ListenAndServe(":8080", nil))
}

func (app *application) Send(w http.ResponseWriter, r *http.Request) {
	var msg Request
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "timmy@tysonjenkins.dev")
	m.SetHeader("To", "timmy@tysonjenkins.dev")
	m.SetHeader("Subject", "Subject")
	m.SetBody("text/html", "message")

	app.client.DialAndSend()
	if err := app.client.DialAndSend(m); err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("Message sent! %s\n", msg.Message)))
}
