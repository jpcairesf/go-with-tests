package selection

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := measureResponseTime(a)
	durationB := measureResponseTime(b)

	if durationA < durationB {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	startA := time.Now()
	_, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	durationA := time.Since(startA)
	return durationA
}
