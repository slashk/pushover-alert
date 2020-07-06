package main // (c) 2020 ken pepple (ken@pepple.io)

import (
	"testing"
)

func Test_CredCheckPositive(t *testing.T) {
	var p pushoverNotification
	cfg := make(map[string]string)
	cfg["pushoverRcpt"] = "sdfsdfsdf"
	cfg["pushoverToken"] = "sfsfsdgsdg"
	if p.new(cfg) != nil {
		t.Errorf("expected pushover creds: %v", p.new(cfg))
	}
	cfg["pushoverRcpt"] = ""
	if p.new(cfg) == nil {
		t.Errorf("expected pushover error: %v", p.new(cfg))
	}
	cfg["pushoverRcpt"] = "sdfsdfsdf"
	cfg["pushoverToken"] = ""
	if p.new(cfg) == nil {
		t.Errorf("expected pushover creds: %v", p.new(cfg))
	}
}
