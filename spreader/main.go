package main

import "fmt"

type total int

func sum(nums ...int) total {
	var t total
	for _, n := range nums {
		t += total(n)
	}
	return t
}
func (t total) String() string {
	if t == 0 {
		return "0"
	}
	return fmt.Sprintf("The fmt will prioritize this method by printing this line == %d \n", t)
}
func main() {
	fmt.Println()
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(sum())
	fmt.Println()

}
