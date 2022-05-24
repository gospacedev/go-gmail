package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(sender *widget.Entry, receiver *widget.Entry, subject *widget.Entry, body *widget.Entry, password *widget.Entry, status *widget.Label) {
	msg := gomail.NewMessage()

	msg.SetHeader("From", sender.Text)
	msg.SetHeader("To", receiver.Text)
	msg.SetHeader("Subject", subject.Text)
	msg.SetBody("text/plain", body.Text)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, sender.Text, password.Text)

	err := dialer.DialAndSend(msg)
	if err != nil {
		fmt.Println(err)

		status.Text = "Invalid Entry"
		status.Refresh()
	} else {
		fmt.Println("Email sent successfully!")

		status.Text = "Email sent successfully!"
		status.Refresh()
	}
}
