package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestActionHandlerURLChallenge(t *testing.T) {
	w := handleChallengeRequest(challengePayload{
		Challenge: "it is a challenge",
		Token:     "it is a verification token",
		Type:      ActionTypeURLVerification,
	})

	if http.StatusOK != w.Code {
		t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
	}

	ct := w.Header().Get("Content-Type")
	if !strings.HasPrefix(ct, "text/plain") {
		t.Fatalf("Expected %s, got %s", "text/plain", ct)
	}

	b, _ := ioutil.ReadAll(w.Body)
	if "it is a challenge" != string(b) {
		t.Fatalf("Expected %s, got %s", "it is a challenge", b)
	}
}

func TestActionHandlerVerificationTokenMismatch(t *testing.T) {
	w := handleChallengeRequest(challengePayload{
		Token: "it is NOT a verification token",
	})

	if http.StatusBadRequest != w.Code {
		t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestActionHandlerUnknownType(t *testing.T) {
	w := handleChallengeRequest(challengePayload{
		Token: "it is a verification token",
		Type:  "unknown_type",
	})

	if http.StatusBadRequest != w.Code {
		t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
	}
}

type challengePayload struct {
	Token     string `json:"token"`
	Type      string `json:"type"`
	Challenge string `json:"challenge"`
}

func handleChallengeRequest(p challengePayload) *httptest.ResponseRecorder {
	c := Config{VerificationToken: "it is a verification token"}
	h := ActionHandler{Config: c}

	w := httptest.NewRecorder()

	m := "POST"
	u := "https://ebooks.thomasd.se/slack/action-endpoint"
	b, _ := json.Marshal(p)
	r := httptest.NewRequest(m, u, bytes.NewReader(b))

	h.ServeHTTP(w, r)

	return w
}
