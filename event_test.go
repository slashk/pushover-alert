package main // (c) 2020 ken pepple (ken@pepple.io)

import (
	"testing"
)

var normalEvent = map[string]string{
	"msg":    "",
	"name":   "pull_request",
	"sha":    "00573ba",
	"ref":    "v1.0.0",
	"repo":   "slashk/pushover-alert",
	"status": "success",
}

var failedEvent = map[string]string{
	"msg":    "",
	"name":   "pull_request",
	"sha":    "00573ba",
	"ref":    "v1.0.0",
	"repo":   "slashk/pushover-alert",
	"status": "failure",
}

var overrideEvent = map[string]string{
	"msg":    "This is a override message",
	"name":   "pull_request",
	"sha":    "00573ba",
	"ref":    "v1.0.0",
	"repo":   "slashk/pushover-alert",
	"status": "success",
}

func Test_NormalMsg(t *testing.T) {
	m := createMsg(normalEvent)
	expected := "👍 <b>Pull Request</b> v1.0.0 from slashk/pushover-alert status: success"
	if m != expected {
		t.Errorf("expected: %v, got: %v", expected, m)
	}
}

func Test_Override(t *testing.T) {
	m := createMsg(overrideEvent)
	expected := overrideEvent["msg"]
	if m != expected {
		t.Errorf("expected: %v, got: %v", expected, m)
	}
}

func Test_FailureEvent(t *testing.T) {
	m := createMsg(failedEvent)
	expected := "👎 <b>Pull Request</b> v1.0.0 from slashk/pushover-alert status: failure"
	if m != expected {
		t.Errorf("expected: %v, got: %v", expected, m)
	}
}
