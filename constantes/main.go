package main

import "fmt"

const (
	a = 1
	b
	c
)
const (
	A = iota // 0
	B        // 1
	C        // 2
)
const (
	X = iota // 0
	_        // Skipped
	Y        // 2
	Z        // 3
)
const (
	KB = 1 << (10 * iota) // 1 << 0 = 1 (1 KB)
	MB                    // 1 << 10 = 1024 (1 MB)
	GB                    // 1 << 20 = 1048576 (1 GB)
)
const (
	FlagA = 1 << iota // 1 (0001)
	FlagB             // 2 (0010)
	FlagC             // 4 (0100)
	FlagD             // 8 (1000)
)
const (
	Red, Green, Blue     = iota, iota, iota // All are 0
	Yellow, Black, White                    // All are 1
)
const (
	Apple  = iota // 0
	Banana        // 1
	Cherry = 1    // Reset to 1
	Durian        // 1
	Mango         // 1
)
const (
	A1 = iota // 0
	B1 = 2    // Explicitly set
	C1 = 2    // Explicitly set
	D1 = iota // 3 (resumes iota sequence)
	E1 = iota // 4
	F1        // 5
)
const Num = 42 // Typeless constant

func main() {
	fmt.Println(a, b, c)                              // outputs 1,1,1
	fmt.Println(A, B, C)                              // outputs 0,1,2
	fmt.Println(X, Y, Z)                              // outputs 0,2,3
	fmt.Println(KB, MB, GB)                           // outputs 1,1024,1048576
	fmt.Println(FlagA, FlagB, FlagC, FlagD)           // outputs 1,2,4,8
	fmt.Println(Red, Green, Blue)                     // outputs 0,0,0
	fmt.Println(Yellow, Black, White)                 // outputs 1,1,1
	fmt.Println(Apple, Banana, Cherry, Durian, Mango) // outputs 0,1,1,1,1
	fmt.Println(A1, B1, C1, D1, E1, F1)
	var i int = Num     // Num is treated as an int
	var f float64 = Num // Num is treated as a float64
	fmt.Println(i, f)   // 42, 42
}
