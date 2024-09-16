package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v: %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

func main() {
	//var e error
	e := &MyError{time.Now(), "Something is broken, program failed"}
	fmt.Println(e)
	fmt.Println(run())
}
