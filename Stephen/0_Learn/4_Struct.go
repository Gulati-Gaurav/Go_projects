package main

import "fmt"

type temp struct {
	name string
	obj  temp2
}
type temp2 struct {
	name2 string
}
func main() {
	firstObj := temp{
		name: "Gaurav",
		obj: temp2{
			name2: "Gulati",
		},
	}
	secondObj := firstObj
	secondObj.obj.name2 = "HINL"
	fmt.Println(firstObj.obj.name2) // Gulati
}

// In Go, structs are value types by default.
// This creates a copy of firstObj. Both secondObj and firstObj are now independent structs in memory. Modifying secondObj does not affect firstObj.
// Nested struct is also copied: The obj field, which is of type temp2, is also copied. So, secondObj.obj is a separate copy of firstObj.obj.
