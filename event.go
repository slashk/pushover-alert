package main

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

	if c["override"] != "" {
		return c["override"]
	}

	m := fmt.Sprintf("%s %s %s from %s was %s",
		icon[c["status"]], fixName(c["name"]), c["ref"], c["repo"], c["status"])
	return m
}

// fixName restyles the github event name into proper english
func fixName(s string) string {
	return strings.Title(strings.ReplaceAll(s, "_", " "))
}
