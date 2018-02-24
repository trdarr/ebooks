package slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCommandHandlerEbooksCommand(t *testing.T) {
	v := url.Values{}
	v.Set("token", "it is a verification token")
	v.Set("command", "/ebooks")

	w := handleCommand(v)

	if http.StatusOK != w.Code {
		t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
	}

	b, _ := ioutil.ReadAll(w.Body)
	if "" != string(b) {
		t.Fatalf("Unexpected empty body, got %s", string(b))
	}
}

func TestCommandHandlerVerificationTokenMismatch(t *testing.T) {
	v := url.Values{}
	v.Set("token", "it is NOT a verification token")

	w := handleCommand(v)

	if http.StatusBadRequest != w.Code {
		t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCommandHandlerUnknownCommand(t *testing.T) {
	v := url.Values{}
	v.Set("token", "it is a verification token")
	v.Set("command", "/unknown")

	w := handleCommand(v)

	if http.StatusOK != w.Code {
		t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
	}

	ct := w.Header().Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		t.Fatalf("Expected %s, got %s", "application/json", ct)
	}

	var r Response
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(b, &r)

	if ResponseTypeEphemeral != r.Type {
		t.Fatalf("Expected %s, got %s", ResponseTypeEphemeral, r.Type)
	}

	if !strings.HasPrefix(r.Text, "Unexpected command") {
		t.Fatalf("Unexpected response %s", r.Text)
	}
}

func handleCommand(v url.Values) *httptest.ResponseRecorder {
	c := Config{VerificationToken: "it is a verification token"}
	h := CommandHandler{Config: c}

	w := httptest.NewRecorder()

	m := "POST"
	u := "https://ebooks.thomasd.se/slack/command"
	r := httptest.NewRequest(m, u, strings.NewReader(v.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	h.ServeHTTP(w, r)

	return w
}
