package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		_, err := fmt.Fprintln(out, i)
		if err != nil {
			return
		}
	}
	_, err := fmt.Fprint(out, "Go!")
	if err != nil {
		return
	}
}
