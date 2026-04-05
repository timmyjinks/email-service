package main

import (
	"github.com/resend/resend-go/v3"
	"log"
)

type ResendClient struct {
	cli *resend.Client
}

func NewClient(key string) *ResendClient {
	return &ResendClient{
		cli: resend.NewClient(key),
	}

}

func (r *ResendClient) SendEmail() {
	params := &resend.SendEmailRequest{
		From:    "TYSONCLOUD <notifications@tysonjenkins.dev>",
		To:      []string{"tyson.j.jenkins@gmail.com"},
		Html:    "<h1>HELLO</h1>",
		Subject: "Hello from Go",
	}

	sent, err := r.cli.Emails.Send(params)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(sent)
}
