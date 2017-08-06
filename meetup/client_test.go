package meetup_test

import (
	"net/http"
	"testing"

	"github.com/eggya/meetup-api/meetup"
)

func TestNewClient(t *testing.T) {
	c, err := meetup.NewClient("123abc")
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if c.URLBase.String() != "https://api.meetup.com" {
		t.Errorf("expected key %s, got: %s", "https://api.meetup.com", c.URLBase)
	}
	if c.HttpClient != http.DefaultClient {
		t.Errorf("expected key %v, got: %v", http.DefaultClient, c.HttpClient)
	}
}
