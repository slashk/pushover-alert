package main // (c) 2020 ken pepple (ken@pepple.io)

import (
	"io"
	"strings"

	"github.com/gregdel/pushover"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
)

type pushoverNotification struct {
	msg      string
	title    string
	url      string
	urlTitle string
}

type pushoverCreds struct {
	rcpt  string
	token string
}

func (p pushoverCreds) notify(n pushoverNotification) (string, error) {
	app := pushover.New(p.token)
	recipient := pushover.NewRecipient(p.rcpt)

	// TODO	set input for html and device
	message := &pushover.Message{
		Message:  n.msg,
		Title:    n.title,
		Priority: pushover.PriorityNormal,
		URL:      n.url,
		// URLTitle: "",
	}

	// TODO figure out pictures
	// pic, err := getPic(s.image)
	// if err != nil {
	// 	log.Infof("notication picture error: %s", err)
	// }
	// if err := message.AddAttachment(pic); err != nil {
	// 	log.Infof("attachment error: %s", err)
	// }

	resp, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Infof("notication error: %s", err.Error())
	}
	return resp.String(), nil
}

func getPic(url string) (io.Reader, error) {
	resp, body, errs := gorequest.New().Get(url).End()
	if len(errs) > 0 {
		log.Debugf("picture download error: %v", errs[0])
		return nil, errs[0]
	}
	if resp.Status != "200 OK" {
		log.Debugf("picture download error: %v status code", resp.Status)
		return nil, nil
	}
	return strings.NewReader(body), nil
}
