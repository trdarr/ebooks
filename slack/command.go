package slack

import "net/http"

type Command struct {
	Token       string
	TeamID      string
	TeamDomain  string
	ChannelID   string
	ChannelName string
	UserID      string
	UserName    string
	Command     string
	Text        string
	ResponseURL string
	TriggerID   string
}

func NewCommand(r *http.Request) Command {
	return Command{
		Token:       r.FormValue("token"),
		TeamID:      r.FormValue("team_id"),
		TeamDomain:  r.FormValue("team_domain"),
		ChannelID:   r.FormValue("channel_id"),
		ChannelName: r.FormValue("channel_name"),
		UserID:      r.FormValue("user_id"),
		UserName:    r.FormValue("user_name"),
		Command:     r.FormValue("command"),
		Text:        r.FormValue("text"),
		ResponseURL: r.FormValue("response_url"),
		TriggerID:   r.FormValue("trigger_id"),
	}
}
