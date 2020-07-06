package main // (c) 2020 ken pepple (ken@pepple.io)

import (
	"errors"
	"time"

	"github.com/gregdel/pushover"
	"github.com/sethvargo/go-githubactions"
)

type pushoverNotification struct {
	rcpt     string
	token    string
	msg      string
	title    string
	url      string
	urlTitle string
	device   string
	sound    string
}

func (p pushoverNotification) new(c map[string]string) error {
	if c["pushoverRcpt"] == "" {
		return errors.New("pushover recipient input (pushover_rcpt) not set")
	}
	if c["pushoverToken"] == "" {
		return errors.New("pushover token input (pushover_token) not set")
	}
	// TODO	check recipient validity
	// recipientDetails, err := app.GetRecipientDetails(recipient)
	p = pushoverNotification{
		rcpt:  c["pushoverRcpt"],
		token: c["pushoverToken"],
		msg:   createMsg(c),
		title: createTitle(c),
		url:   createURL(c),
		sound: c["sound"],
	}
	if c["device"] != "" {
		p.device = c["device"]
	}
	return nil
}

func (p pushoverNotification) notify() (string, error) {
	app := pushover.New(p.token)
	recipient := pushover.NewRecipient(p.rcpt)

	message := &pushover.Message{
		Message:   p.msg,
		Title:     p.title,
		Priority:  pushover.PriorityNormal,
		URL:       p.url,
		Sound:     p.sound,
		HTML:      true,
		URLTitle:  p.urlTitle,
		Timestamp: time.Now().Unix(),
	}
	if p.device != "" {
		message.DeviceName = p.device
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
		g := githubactions.New()
		g.Fatalf("notication sending error: %s", err.Error())
	}
	return resp.String(), nil
}

// func getPic(url string) (io.Reader, error) {
// 	resp, body, errs := gorequest.New().Get(url).End()
// 	if len(errs) > 0 {
// 		log.Debugf("picture download error: %v", errs[0])
// 		return nil, errs[0]
// 	}
// 	if resp.Status != "200 OK" {
// 		log.Debugf("picture download error: %v status code", resp.Status)
// 		return nil, nil
// 	}
// 	return strings.NewReader(body), nil
// }
