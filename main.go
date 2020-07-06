package main // (c) 2020 ken pepple (ken@pepple.io)

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

	ga := githubactions.New()
	config["pushoverRcpt"] = ga.GetInput("pushover_rcpt")   // os.Getenv("INPUT_PUSHOVER_RCPT")
	config["pushoverToken"] = ga.GetInput("pushover_token") // os.Getenv("INPUT_PUSHOVER_TOKEN")
	config["msg"] = ga.GetInput("msg")                      // os.Getenv("INPUT_MSG")
	config["device"] = ga.GetInput("device")                // os.Getenv("INPUT_DEVICE")
	config["title"] = ga.GetInput("title")                  // os.Getenv("INPUT_TITLE")
	config["priority"] = ga.GetInput("priority")            // os.Getenv("INPUT_PRIORITY")
	config["sound"] = ga.GetInput("sound")                  // os.Getenv("INPUT_SOUND")
}

func main() {
	var p pushoverNotification

	g := githubactions.New()
	g.Debugf("%v, commit %v, built at %v\n", version, commit, date)
	g.AddMask(config["rcpt"])
	g.AddMask(config["token"])
	g.Debugf("Configs set: %v", config)

	err := p.new(config)
	if err != nil {
		g.Fatalf("Error with pushover credentials: %s", err)
	}
	g.Debugf("message: %v", p)

	// send notification to device
	o, err := p.notify()
	if err != nil {
		g.Fatalf("notification failed: %v", err)
	}
	g.Debugf("notification successful: %v", o)
}
