package main

import (
	"fmt"

	"github.com/robfig/config"
	"github.com/sendgrid/sendgrid-go"
)

func main() {
	c, _ := config.ReadDefault("mail.conf")
	user, _ := c.String("sendgrid", "user")
	key, _ := c.String("sendgrid", "key")
	to, _ := c.String("sendgrid", "to")
	name, _ := c.String("sendgrid", "to_name")
	from, _ := c.String("sendgrid", "from")
	subject, _ := c.String("sendgrid", "subject")
	body, _ := c.String("sendgrid", "body")

	sg := sendgrid.NewSendGridClient(user, key)
	message := sendgrid.NewMail()
	message.AddTo(to)
	message.AddToName(name)
	message.SetSubject(subject)
	message.SetText(body)
	message.SetFrom(from)
	if r := sg.Send(message); r == nil {
		fmt.Println("sent")
	} else {
		fmt.Println(r)
	}
}
