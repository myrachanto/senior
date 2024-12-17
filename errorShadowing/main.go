package main

import (
	"errors"
	"fmt"
)
func doSomething() error {
	return errors.New("something went wrong")
}
func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Error in main:", err)
	}
	// Error shadowing here!
	if err := doSomething(); err != nil {
		// This `err` shadows the outer `err`, so any logic after this
		// block that checks `err` might get confusing results.
		fmt.Println("Error in block:", err)
	}
	// `err` here is still the one from the first call to doSomething.
	if err != nil {
		fmt.Println("Error after block:", err)
	}
}
