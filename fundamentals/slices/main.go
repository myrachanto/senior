package main

import "fmt"

func main() {
	// Create an array of strings with four elements
	arr := [5]string{"one", "two", "three", "four"}
	// Create a slice from the array with arr[2:5]
	slice := arr[2:5]
	// Print the slice, length, and capacity)
	fmt.Printf("slice = %v cap %d len = %d \n", slice, cap(slice), len(slice))
	slice = append(slice, "six")
	fmt.Printf("slice = %v cap %d len = %d\n", slice, cap(slice), len(slice))
	changer(slice)
	fmt.Printf("slice = %v cap %d len = %d\n", slice, cap(slice), len(slice))
	slice2 := arr[:]
	fmt.Printf("slice2 = %v cap %d len = %d\n", slice2, cap(slice2), len(slice2))
	slice3 := slice2[1:3]
	fmt.Printf("slice3 = %v cap %d len = %d\n", slice3, cap(slice3), len(slice3))
	slice4 := slice3[2:4]
	fmt.Printf("slice4 = %v cap %d len = %d\n", slice4, cap(slice4), len(slice4))
}
func changer(data []string) {
	data[1] = "1"
	fmt.Printf("slice = %v cap %d len = %d\n", data, cap(data), len(data))
}
