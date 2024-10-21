package selection

import (
	"fmt"
	"net/http"
	"time"
)

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		close(ch)
	}()
	return ch
}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
