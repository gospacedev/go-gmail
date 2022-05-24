package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()

	win := app.NewWindow("Go Gmail")

	win.Resize(fyne.NewSize(500, 400))

	win.SetIcon(resourceIconPng)

	sender := widget.NewEntry()
	sender.SetPlaceHolder("From:")

	password := widget.NewEntry()
	password.SetPlaceHolder("Password:")

	receiver := widget.NewEntry()
	receiver.SetPlaceHolder("To:")

	subject := widget.NewEntry()
	subject.SetPlaceHolder("Subject:")

	body := widget.NewMultiLineEntry()
	body.SetPlaceHolder("Message:")

	status := widget.NewLabel("Please enter all fields...")

	SendButton := widget.NewButton("Send Email", func() {
		SendEmail(sender, receiver, subject, body, password, status)
	})

	win.SetContent(container.NewVBox(
		sender,
		password,
		receiver,
		subject,
		body,
		SendButton,
		status,
	))

	win.ShowAndRun()
}
