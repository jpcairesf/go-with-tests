package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		_, err := fmt.Fprintln(out, i)
		if err != nil {
			return
		}
		sleeper.Sleep()
	}
	_, err := fmt.Fprint(out, "Go!")
	if err != nil {
		return
	}
}
