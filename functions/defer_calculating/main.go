package main

import "fmt"

// Need to show solution

const (
	StatusOk    = "ok"
	StatusError = "error"
)

func notify(status string) {
	fmt.Println(status)
}

func process() {
	var status string
	defer func() {
		notify(status)
	}()

	// processing..
	status = StatusError
}

func main() {
	process()
}
