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
	resend *ResendClient
}

type Request struct {
	Message string `json:"message"`
}

func main() {
	cfg := Load()

	d := gomail.NewDialer("mail.tysonjenkins.dev", 587, "timmy@tysonjenkins.dev", "password")
	d.TLSConfig = &tls.Config{ServerName: "mail.tysonjenkins.dev", InsecureSkipVerify: true}
	d.LocalName = "mail.tysonjenkins.dev"

	app := application{
		client: d,
		resend: NewClient(cfg.Resend.KEY),
	}

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
	app.resend.SendEmail()

	w.Write([]byte(fmt.Sprintf("Message sent! %s\n", msg.Message)))
}
