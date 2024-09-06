package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"test-pari/schemas"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type SMTP struct {
	Config *schemas.SMTPConfig
}

func InitEmail(config *schemas.SMTPConfig) *SMTP {
	return &SMTP{
		Config: config,
	}
}

func (e *SMTP) Send(to []string, cc []string, bcc []string, subject string, bodyType string, body string, attachment []string) error {
	log.Println("Email - Send - starting...")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.Config.Email)
	mailer.SetHeader("To", to...)

	if cc != nil && len(cc) > 0 {
		mailer.SetHeader("Cc", cc...)
	}

	if bcc != nil && len(bcc) > 0 {
		mailer.SetHeader("Bcc", bcc...)
	}

	mailer.SetHeader("Subject", subject)
	mailer.SetBody(bodyType, body)

	if attachment != nil && len(attachment) > 0 {
		for _, path := range attachment {
			mailer.Attach(path)
		}
	}

	dialer := gomail.NewDialer(
		e.Config.Host,
		e.Config.Port,
		e.Config.Email,
		e.Config.Password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	log.Println("Email - Send - finished")
	return nil
}

func (e *SMTP) SendThirdPartyAPI(to string, subject string, body string) error {

	req := schemas.PushEmail{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	jsonReq, err := json.Marshal(req)
	resp, err := http.Post("http://contoh/aja", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var cs schemas.ResponseExample
	err = json.Unmarshal(bodyBytes, &cs)
	if err != nil {
		return err
	}

	if cs.Status != "success" {
		return fmt.Errorf(cs.Code)
	}
	return nil
}

func GenerateOTPLayout(code string, name string, operation string) (res string, err error) {
	h := hermes.Hermes{

		Product: hermes.Product{
			Name:      "San Diego Hills",
			Link:      "https://www.sandiegohills.co.id/",
			Logo:      "https://sandiegohills.co.id/support/images/logoPrintln.png",
			Copyright: "© 2015 SANDIEGOHILLS.CO.ID. ALL RIGHTS RESERVED",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"You are requesting " + operation,
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please copy your otp code :",
					InviteCode:   code,
				},
			},
			Outros: []string{
				"Thank you for your trust to san diego hills",
			},
		},
	}

	return h.GenerateHTML(email)
}
