package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/robfig/config"
	"github.com/sendgrid/sendgrid-go"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n %s [options] ip\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	config_file := flag.String("config", "mail.conf", "path to the config file")
	flag.Parse()
	c, _ := config.ReadDefault(*config_file)
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
