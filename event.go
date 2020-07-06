package main // (c) 2020 ken pepple (ken@pepple.io)

import (
	"fmt"
	"strings"
)

func createMsg(c map[string]string) string {
	icon := map[string]string{
		"success":   "👍",
		"failure":   "👎",
		"cancelled": "🤚",
	}
	if c["msg"] != "" {
		return c["msg"]
	}
	m := fmt.Sprintf("%s %s %s from %s status: %s",
		icon[c["status"]], fixName(c["name"]), c["ref"], c["repo"], c["status"])
	return m
}

func createTitle(c map[string]string) string {
	if c["title"] != "" {
		return c["title"]
	}
	return fmt.Sprintf("%s %s on %s", fixName(c["name"]), c["ref"], c["repo"])
}

func createURL(c map[string]string) string {
	return "https://github.com/" + c["repo"]
}

// fixName restyles the github event name into proper english
func fixName(s string) string {
	return strings.Title(strings.ReplaceAll(s, "_", " "))
}
