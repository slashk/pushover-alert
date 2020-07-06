package main // (c) 2020 ken pepple (ken@pepple.io)

import (
	"testing"
)

func Test_CredCheckPositive(t *testing.T) {
	var p pushoverNotification
	var cfg map[string]string
	cfg["rcpt"] = "sdfsdfsdf"
	cfg["token"] = "sfsfsdgsdg"
	if p.new(cfg) != nil {
		t.Errorf("expected pushover creds: %v", p.new(cfg))
	}
	cfg["rcpt"] = ""
	if p.new(cfg) == nil {
		t.Errorf("expected pushover error: %v", p.new(cfg))
	}
	cfg["rcpt"] = "sdfsdfsdf"
	cfg["token"] = ""
	if p.new(cfg) == nil {
		t.Errorf("expected pushover creds: %v", p.new(cfg))
	}
}
