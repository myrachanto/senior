package main

import "fmt"

// be very careful of how you use defer with anonymous functions
func checker() {
	var num int
	defer fmt.Println(num)
	num++
	defer func() {
		num++
		fmt.Println(num)
	}()
	defer func() {
		num++
		fmt.Println(num)
	}()
}

func main() {
	checker()
	// output 2, 3, 0
}
