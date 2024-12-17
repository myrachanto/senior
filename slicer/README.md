What are slices and how do they work in Golang?


Hi everyone 🤗 , let's dive into slices basics and fill in the gaps!


✳️ A slice in Go (Golang) is a flexible, dynamically sized view of an array's elements.


1️⃣ Dynamic Size:

✅ Slices can grow or shrink as needed, unlike arrays with a fixed size. 

✅ Their capacity can increase automatically when appending elements. (they double)


2️⃣ Underlying Array:

✅ A slice does not store its data. Instead, it is a descriptor that points to an array where the actual data is stored. (as a result when a slice is passed to a function and modified the the outside scope of the function gets to see the changes)


3️⃣ Structure:

✅ A pointer to the underlying array.

✅ The length (number of elements in the slice).

✅ The capacity (maximum number of elements it can hold before reallocating).


4️⃣ Declaration and Initialization:

✅ `t := make([]int, 5) // Creates a slice of length 5, capacity 5` appending 1 to it will result to [000001]

✅ `t := make([]int, 0,5) // Creates a slice of length 0, capacity 5` appending 1 to it will result to [1]


5️⃣ Slicing an array:

```bash
   arr := [5]int{10, 20, 30, 40, 50}
   

    // Create slices

    s1 := arr[1:4]                    // Slice of elements [20, 30, 40]

    fmt.Println(len(s1), cap(s1), s1) // Output: 3,4, [20 30 40  ]

```

✅ The capacity of a slice is the number of elements in the underlying array starting from the slice's start index, up to the end of the array. calculated as:

`cap(s1) = total_elements_in_array - start_index`

In this case it's 4


✅ The length of a slice is the number of elements in it, calculated as:

`len(s1) = end_index - start_index`

 --- in this case 3


6️⃣ Built-in Functions to work with slices:

✅ len(slice): Returns the current length of the slice.

✅ cap(slice): Returns the capacity of the slice.

✅ append(slice, values...): Adds elements to the slice, potentially increasing its capacity. (by doubling it)


7️⃣ Nil vs Empty Slice:

✅ A nil slice has no underlying array (var s []int), and both len(s) and cap(s) are 0

✅ An empty slice has an underlying array of zero length (s := []int{}), but len(s) is still 0.


✳️ you can append to both nil slices and empty slices. After appending, both will behave similarly, as they gain an underlying array to store the appended elements. The key difference lies in their initial states and their comparison with nil

#go #golangtips #golang #masteringgolang #backeddeveloper #backend