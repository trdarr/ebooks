package slack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CommandHandler struct {
	Config
}

func (h CommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error parsing form: %s", err)
		return
	}

	c := NewCommand(r)

	if c.Token != h.Config.VerificationToken {
		s := "Verification token doesn’t match"
		http.Error(w, s, http.StatusBadRequest)
		log.Print(s)
		return
	}

	log.Printf(
		"Got command: user %s (%s), channel %s (%s), team %s (%s): %s",
		c.UserName, c.UserID,
		c.ChannelName, c.ChannelID,
		c.TeamDomain, c.TeamID,
		c.Command,
	)

	switch c.Command {
	case "/ebooks":
		r, err := json.Marshal(Response{
			Text: "ebooks",
			Type: ResponseTypeInChannel,
		})

		if err != nil {
			s := fmt.Sprintf("Error marshalling response: %s", err)
			http.Error(w, s, http.StatusInternalServerError)
			log.Print(s)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, string(r))

	default:
		s := fmt.Sprintf("Unexpected command: %s", c.Command)
		log.Print(s)

		r, err := json.Marshal(Response{
			Text: s,
			Type: ResponseTypeEphemeral,
		})

		if err != nil {
			s = fmt.Sprintf("Error marshalling response: %s", err)
			http.Error(w, s, http.StatusInternalServerError)
			log.Print(s)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, string(r))
	}
}
