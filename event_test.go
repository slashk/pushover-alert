package main

import (
	"testing"
)

var normalEvent = map[string]string{
	"override": "",
	"name":     "pull_request",
	"sha":      "00573ba",
	"ref":      "v1.0.0",
	"repo":     "slashk/pushover-alert",
	"status":   "success",
}

var failedEvent = map[string]string{
	"override": "",
	"name":     "pull_request",
	"sha":      "00573ba",
	"ref":      "v1.0.0",
	"repo":     "slashk/pushover-alert",
	"status":   "failure",
}

var overrideEvent = map[string]string{
	"override": "This is a override message",
	"name":     "pull_request",
	"sha":      "00573ba",
	"ref":      "v1.0.0",
	"repo":     "slashk/pushover-alert",
	"status":   "success",
}

func Test_NormalMsg(t *testing.T) {
	m := createMsg(normalEvent)
	expected := "👍 Pull Request v1.0.0 from slashk/pushover-alert was success"
	if m != expected {
		t.Errorf("expected: %v, got: %v", expected, m)
	}
}

func Test_Override(t *testing.T) {
	m := createMsg(overrideEvent)
	expected := overrideEvent["override"]
	if m != expected {
		t.Errorf("expected: %v, got: %v", expected, m)
	}
}

func Test_FailureEvent(t *testing.T) {
	m := createMsg(failedEvent)
	expected := "👎 Pull Request v1.0.0 from slashk/pushover-alert was failure"
	if m != expected {
		t.Errorf("expected: %v, got: %v", expected, m)
	}
}
