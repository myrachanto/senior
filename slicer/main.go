package main

import "fmt"

// A slice is represented internally by:
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {
	var s []int             // Declares a nil slice
	s = []int{1, 2, 3}      // Initializes a slice with values
	t := make([]int, 5)     // Creates a slice of length 5, capacity 5
	u := make([]int, 5, 10) // Creates a slice of length 5, capacity 10

	fmt.Println(s, t, u) // outputs [1 2 3] [0 0 0 0 0] [0 0 0 0 0]

	// Create an array
	arr := [5]int{10, 20, 30, 40, 50}
	// Create slices
	s1 := arr[1:4]                    // Slice of elements [20, 30, 40]
	fmt.Println(len(s1), cap(s1), s1) // Output: 3,4, [20 30 40 ]

	// Append to the slice
	s1 = append(s1, 60)
	fmt.Println(len(s1), cap(s1), s1) // Output: 4,4, [20 30 40 60]
	s1 = append(s1, 70)
	fmt.Println(len(s1), cap(s1), s1) // Output:5,8 [20 30 40 60,70]

}
