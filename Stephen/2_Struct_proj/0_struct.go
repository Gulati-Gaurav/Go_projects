package main

// In Go, you can create a type with multiple data members by defining a struct.

// In Go, whether a struct is stored on the heap or stack depends on the context in which it is used. The Go compiler uses escape analysis to decide whether a variable should be allocated on the stack or the heap. Here are the general rules:
// Stack Allocation: If the compiler determines that a variable's lifetime is confined to the scope in which it is defined (e.g., within a function), it will allocate that variable on the stack. This includes struct instances that do not escape the function.
// Heap Allocation: If the compiler determines that a variable's lifetime extends beyond the scope in which it is defined (e.g., the variable is returned from a function or referenced by a goroutine), it will allocate that variable on the heap.

// If you want to see the escape analysis details, you can compile your Go code with the -gcflags=-m flag, which will show where allocations are made:
// go build -gcflags=-m main.go
// This command will output information about whether variables are escaping to the heap or remaining on the stack.

// No inheritance in struct

func main() {

	Initialisation()

	Receiver()

	// to create temporary structs
	temp_structs_go()

	// structs are value type
	tempFunc()

	// Deep copy a slice using either the built-in copy or append function to get a duplicated slice. However, assign an existed slice to a new variable gets a shallow copy whose slice header points to the same backing array as the source slice.

	// One important thing regarding pointers. Say we pass a slice by value to receiver func then original slice will be changed. The reason is slice is pointer to underlying to array hence changes.

	// When we create a slice Go creates a slice and array data structure for us. Slice has variables ptrToHead, length, capacity. ptrToHead points to the array in the memory. When we call slice we don't refer to slice but array. Go is still pass by value. Hence it creates a new slice for us having 3 new variabes ptrToHead, length, capacity. Values are copied hence ptrToHead is also copied. But ptrToHead still points to the same original array hence no change.

	// Not only slice but other reference types behaves the same way. Slice is reference type variable. Other includes => map, channels, pointers, functions.
	// Value types include => int, float, string, bool, structs. Use pointer to change these things in the function.

	// If you pass a struct by value to a function, a copy of the struct is made, regardless of whether it resides on the heap or stack.

	// Below also works
}
type MyStruct struct {
	value int
}

func (ms MyStruct) changeNum() {
	ms.value = 5
}

func createStruct() *MyStruct {
	s := MyStruct{value: 42}
	// s is allocated on the heap because its reference is returned
	return &s
}

func main() {
	ptr := createStruct()
	// ptr is a pointer to the heap-allocated MyStruct
	fmt.Println(ptr.value) // Output: 42
	fmt.Println((*ptr).value) // Output: 42 see here both above and this lines works

ptr.changeNum()
fmt.Println(ptr.value)
}
