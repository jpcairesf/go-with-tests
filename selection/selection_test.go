package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := newDelayedServer(20 * time.Millisecond)
	slowURL := slowServer.URL
	defer slowServer.Close()

	fastServer := newDelayedServer(0 * time.Millisecond)
	fastURL := fastServer.URL
	defer fastServer.Close()

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func newDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, _ *http.Request) {
			time.Sleep(delay)
			writer.WriteHeader(http.StatusOK)
		}))
	return slowServer
}
