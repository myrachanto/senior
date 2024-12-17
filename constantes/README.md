Understanding constants and enums in Golang #quizthursday.


Some stuff can only be learned or taught or experienced 


Hi everyone 🤗 , feliz navidad, has been running  around in my head


be that as it may, here are a few lessons about constants and enums(iota)


1️⃣ In Go, if a constant in a group is not explicitly assigned a value, it takes the value of the previous constant.

```bash
const (
	a = 1 // outputs 1
	b     // outputs 1
	c     // outputs 1
)
```

2️⃣ The iota keyword is a powerful tool for generating sequential values in constant blocks. (starts at 0)
```bash
const (
	A = iota // 0
	B        // 1
	C        // 2
)

```

3️⃣ Use _ to skip certain values if needed in a constant block.

```bash

const (
	X = iota // 0
	_        // Skipped
	Y        // 2
	Z        // 3
)
```

4️⃣ multiple constants with the same iota value in one line have the same value in a constant block.
```bash
const (
	Red, Green, Blue     = iota, iota, iota // All are 0
	Yellow, Black, White                    // All are 1
)

```


5️⃣ If you set a value in an enum sequence, everything after it will inherit that set value unless you explicitly use iota again or assign a new value in a constant block.

```bash
const (
	Apple  = iota // 0
	Banana        // 1
	Cherry = 1    // Reset to 1
	Durian        // 1
	Mango         // 1
)

```


6️⃣ Each const block has its own iota sequence, and they are independent.


7️⃣ Constants in Go are typeless until explicitly assigned a type. Typeless constants can take on different types depending on the context in which they are used.

```bash

const Num = 42 // Typeless constant
func tester(){
	var i int = Num     // Num is treated as an int
	var f float64 = Num // Num is treated as a float64
	fmt.Println(i, f)   // 42, 42
}
```

#go #golang #golangtips. #senior #backend #softwareengineer