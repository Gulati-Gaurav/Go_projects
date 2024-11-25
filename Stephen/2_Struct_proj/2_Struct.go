package main

import "fmt"

type temp struct {
	name   string
	slices []int
	obj    temp2
}
type temp2 struct {
	name2 string
}

// In Go, structs are value types by default.
// This creates a copy of firstObj. Both secondObj and firstObj are now independent structs in memory. Modifying secondObj does not affect firstObj.
// Nested struct is also copied: The obj field, which is of type temp2, is also copied. So, secondObj.obj is a separate copy of firstObj.obj.


func tempFunc() {
	tempFunc1()
	tempFunc2()
	tempFunc3()
}

// shallow copy proof
func tempFunc1() {
	firstObj := temp{
		name: "Gaurav",
		obj: temp2{
			name2: "Gulati",
		},
	}
	firstObj.slices = make([]int, 4, 5) // to mention len and cap (first len)
	firstObj.slices[0] = 1
	firstObj.slices[1] = 2
	firstObj.slices[2] = 3

	secondObj := firstObj
	secondObj.obj.name2 = "HINL"
	fmt.Println(firstObj.obj.name2) // Gulati

	// Below is deep copy
	secondObj.slices[3] = 4
	fmt.Println(firstObj.slices) // [1, 2, 3, 4]
}

// shallow copy example but seems like deep copy
func tempFunc2() {
	firstObj := temp{
		name:   "Gaurav",
		slices: []int{1, 2},
		obj: temp2{
			name2: "Gulati",
		},
	}

	secondObj := firstObj

	secondObj.slices = append(secondObj.slices, 3)
	fmt.Println(firstObj.slices) // [1, 2]

	// reason is that initially firstObj.slices and secondObj.slices point to same underlying array. 
	// both have capacity of 2
	// when we append to the secondObj slice it sees that capacity is not enough, hence allocates new underlying array and hence the secondObj slice and firstObj slice become detached and hence firstObj doesn't reflect the changes
}

// shallow copy example but seems like deep copy
func tempFunc3() {
	firstObj := temp{
		name: "Gaurav",
		slices: []int{1, 2},
		obj: temp2{
			name2: "Gulati",
		},
	}

	secondObj := firstObj

	secondObj.slices = []int{4, 5}
	fmt.Println(firstObj.slices) // [1, 2]

	// Below when you modify secondObj.slices = []int{4, 5}, you create a new slice with a new underlying array. This breaks the connection between firstObj.slices and secondObj.slices.
}
