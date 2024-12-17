package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range 10 {
		go func() {
			fmt.Println(i, "===>", i*2)
		}()
	}
	time.Sleep(1 * time.Second)
}
