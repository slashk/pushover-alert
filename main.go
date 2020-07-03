package main

import (
	"os"

	"github.com/sethvargo/go-githubactions"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var config map[string]string

func init() {
	config = make(map[string]string)

	// set Github action event variables
	config["eventName"] = os.Getenv("GITHUB_EVENT_NAME")
	config["sha"] = os.Getenv("GITHUB_SHA")
	config["ref"] = os.Getenv("GITHUB_REF")
	config["repo"] = os.Getenv("GITHUB_REPOSITORY")
	config["status"] = os.Getenv("JOB_STATUS")
}

func getInputOrDie(g githubactions.Action, i string) string {
	x := g.GetInput(i)
	if x == "" {
		g.Fatalf("Missing input: " + i)
	}
	return x
}

func main() {
	ga := githubactions.New()

	// set input/output settings from action
	// we can't set this in init as it breaks testings
	config["pushoverRcpt"] = getInputOrDie(*ga, "pushover_rcpt")   // os.Getenv("INPUT_PUSHOVER_RCPT")
	config["pushoverToken"] = getInputOrDie(*ga, "pushover_token") // os.Getenv("INPUT_PUSHOVER_TOKEN")
	config["override"] = ga.GetInput("pushover_override")          // os.Getenv("INPUT_OVERRIDE_MSG")
	p := pushoverCreds{
		rcpt:  config["pushoverRcpt"],
		token: config["pushoverToken"],
	}
	ga.Debugf("Configs set: %v", config)

	// compose alert according to event type
	notification := pushoverNotification{
		msg:   createMsg(config),
		title: "Title",
		url:   "https://github.com",
	}
	ga.Debugf("message: %s", notification)

	// send notification to device
	o, err := p.notify(notification)
	if err != nil {
		ga.Fatalf("notification failed: %v", err)
	}
	ga.Debugf("notification successful: %v", o)
}
