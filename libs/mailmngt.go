package libs

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
)

type DraperMail struct {
}

func (DraperMail) SendMail(toMail string, subject string, mailbody string) {
	// Choose auth method and set it up
	auth := LoginAuth("mustafasahinim28@gmail.com", "5714.Ms98")
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{toMail}
	msg := []byte(fmt.Sprintf("To: %s\n%s\n\n\r\n", toMail, mimeHeaders) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		fmt.Sprintf("%s\r\n", mailbody))
	err := smtp.SendMail("smtp.gmail.com:587", auth, "from@example.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}
