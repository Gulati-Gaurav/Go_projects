package main

import "fmt"

type temp struct {
	name string
	slices []int
	obj  temp2
}
type temp2 struct {
	name2 string
}
func tempFunc() {
	firstObj := temp{
		name: "Gaurav",
		// slices: []int{1, 2},
		obj: temp2{
			name2: "Gulati",
		},
	}
	firstObj.slices = make([]int, 4)
	firstObj.slices[0] = 1
	firstObj.slices[1] = 2
	firstObj.slices[2] = 3
	firstObj.slices = append(firstObj.slices, 3)
	
	secondObj := firstObj
	secondObj.obj.name2 = "HINL"
	fmt.Println(firstObj.obj.name2) // Gulati

	secondObj.slices = append(secondObj.slices, 5)
	fmt.Println(firstObj.slices) // [1, 2]

	// Below when you modify secondObj.slices = []int{4, 5}, you create a new slice with a new underlying array. This breaks the connection between firstObj.slices and secondObj.slices.
	// secondObj.slices = []int{4,5} 
}

// In Go, structs are value types by default.
// This creates a copy of firstObj. Both secondObj and firstObj are now independent structs in memory. Modifying secondObj does not affect firstObj.
// Nested struct is also copied: The obj field, which is of type temp2, is also copied. So, secondObj.obj is a separate copy of firstObj.obj.
