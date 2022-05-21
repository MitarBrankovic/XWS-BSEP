package application

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

type MailService struct {
}

func NewMailService() *MailService {
	return &MailService{}
}

func (service *MailService) SendActivationEmail(token string) {
	/*from := "<xwsclient@gmail.com>"
	password := "<XWSClient123.>"

	toEmailAddress := "<radisamilovcevic3@gmail.com>"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "465" //587
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}*/

	from := mail.Address{"", "xwsclient@gmail.com"}
	to := mail.Address{"", "xwsclient@gmail.com"}
	subj := "Activation mail"
	body := "http://localhost:8000/activate/" + token

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	//servername := "smtp.example.tld:465"
	servername := "smtp.gmail.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", "xwsclient@gmail.com", "XWSClient123.", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}
