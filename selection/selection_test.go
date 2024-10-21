package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one",
		func(t *testing.T) {
			slowServer := newDelayedServer(20 * time.Millisecond)
			slowURL := slowServer.URL
			defer slowServer.Close()

			fastServer := newDelayedServer(0 * time.Millisecond)
			fastURL := fastServer.URL
			defer fastServer.Close()

			want := fastURL
			got, err := ConfigurableRacer(slowURL, fastURL, 1*time.Second)
			if err != nil {
				t.Error(err)
			}

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})

	t.Run("returns an error if a server doesn't respond within 10s",
		func(t *testing.T) {
			serverA := newDelayedServer(12 * time.Millisecond)
			defer serverA.Close()

			serverB := newDelayedServer(11 * time.Millisecond)
			defer serverB.Close()

			_, err := Racer(serverA.URL, serverB.URL)
			if err == nil {
				t.Error("got nil, want error")
			}
		})
}

func newDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(
		func(writer http.ResponseWriter, _ *http.Request) {
			time.Sleep(delay)
			writer.WriteHeader(http.StatusOK)
		}))
	return slowServer
}
