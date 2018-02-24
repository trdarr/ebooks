package main

import (
	"log"
	"net/http"

	"go.thomasd.se/ebooks/slack"
)

func main() {
	config, err := NewConfig("/etc/ebooks/config.json")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/slack/action-endpoint", slack.ActionHandler{Config: config.Slack})
	http.Handle("/slack/command", slack.CommandHandler{Config: config.Slack})
	log.Fatal(http.ListenAndServe(":80", nil))
}
