package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ActionHandler struct {
	Config
}

func (h ActionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error reading request body: %s", err)
		return
	}

	var p struct {
		Token string `json:"token"`
		Type  string `json:"type"`
	}

	if err := json.Unmarshal(b, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error unmarshalling request body: %s", err)
		return
	}

	if p.Token != h.Config.VerificationToken {
		s := "Verification token doesn’t match"
		http.Error(w, s, http.StatusBadRequest)
		log.Print(s)
		return
	}

	switch p.Type {
	case "url_verification":
		h.verifyURL(w, b)

	default:
		s := fmt.Sprintf("Unexpected type: %s", p.Type)
		http.Error(w, s, http.StatusBadRequest)
		log.Print(s)
	}
}

// https://api.slack.com/events/url_verification
func (h ActionHandler) verifyURL(w http.ResponseWriter, body []byte) {
	var p struct {
		Challenge string `json:"challenge"`
	}

	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error unmarshalling request body: %s", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, p.Challenge)
}
