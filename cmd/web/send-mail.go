package main

import (
	// built in Golang packages
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
	// External packages/dependencies
	"github.com/xhit/go-simple-mail/v2"
	// My own packages
	"github.com/darkside1809/bookings/pkg/models"
	
)

// listenForMail creates a goroutine, that wait for sending email from user to user
func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMessage(msg)
		}
	}()
}
// sendMessage creates new smtp server, 
// -connect client to server,
// -read messages from given template/file
// -send email message from given sender address to the given receiver address
func sendMessage(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second


	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}
		mailTemplate := string(data)
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)

	}

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent")
	}
}