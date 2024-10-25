package main

import "fmt"

// In Go, you can create a type with multiple data members by defining a struct.

// In Go, whether a struct is stored on the heap or stack depends on the context in which it is used. The Go compiler uses escape analysis to decide whether a variable should be allocated on the stack or the heap. Here are the general rules:

// Stack Allocation: If the compiler determines that a variable's lifetime is confined to the scope in which it is defined (e.g., within a function), it will allocate that variable on the stack. This includes struct instances that do not escape the function.

// Heap Allocation: If the compiler determines that a variable's lifetime extends beyond the scope in which it is defined (e.g., the variable is returned from a function or referenced by a goroutine), it will allocate that variable on the heap.

// If you want to see the escape analysis details, you can compile your Go code with the -gcflags=-m flag, which will show where allocations are made:
// go build -gcflags=-m main.go
// This command will output information about whether variables are escaping to the heap or remaining on the stack.

// No inheritance in struct

type person struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email   string
	pincode int
}

type person2 struct {
	firstName string
	lastName  string
	contact   contactInfo // See we can have a class as object too
}

type person3 struct {
	firstName   string
	lastName    string
	contactInfo //  imp This is shortcut. means field name is also contactInfo
}

// receiver func for struct
func (p person3) printPerson() {
	fmt.Printf("%+v \n", p)
}

func (p person3) changeFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person3) changeFirstNameRef(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person3) changeFirstNameRef2(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person3) changeFirstNameRef3(newFirstName string) {
	p.firstName = newFirstName
	// Since p is already a pointer, accessing p.firstName directly modifies the underlying struct's field.
	// In essence, when you use a pointer receiver, you're working directly with the original struct, allowing you to modify its fields within the method.
}

func main() {
	// Initialisation_ways
	// 1st instant initialisation
	gaurav := person{"Gaurav", "Gulati"}

	// 2nd instant initialisation
	david := person{firstName: "David", lastName: "Allen"}
	// same as 2nd
	barry := person{
		firstName: "Barry",
		lastName:  "Allen",
	}
	// It is important to put comma in instant initialisation not in struct declaration above
	// Also commas need to placed after all properties including last property

	// 3rd instant initialisation.
	// If we don't put values, struct puts 0 value by default. means empty string for string, 0 for int, float. false for bool
	var bill person
	// to put / change values
	bill.firstName = "Bill"
	bill.lastName = "Gates"

	// Ways to print
	fmt.Println(gaurav, david, bill) // will print properties space separated.
	// If want string formatting or interpolation we use printf instead of println. So that we can use the %v
	// f in Printf or Errorf is used for %v or %+v. %+v will print all property fields and their values(: in b/w the 2)
	// \n means new line
	// see all types of % in https://gobyexample.com/string-formatting
	fmt.Printf("%+v \n", barry)

	// Embedded struct
	john := person2{ // with different name as struct contactInfo
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "johndoe.com",
			pincode: 110034,
		},
	}
	jane := person3{ // with same name as struct contactInfo
		firstName: "Jane",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "janedoe.com",
			pincode: 110053,
		},
	}

	fmt.Printf("%+v \n", john)
	// using receiver func
	jane.printPerson()

	// IMP receiver is passed by value not reference => will not change original object. Hence will have to use pointers
	// IMP Go is a pass by value language whether you send in arguements or as receiver. See what is C#
	jane.changeFirstName("Jain")
	jane.printPerson()

	// passing pointer to make changes in the original object. Basically now we are sending the address of the variable instead of letting create a copy of the value. No problem if go creates a copy of address. See Image
	(&jane).changeFirstNameRef("Jain")
	jane.printPerson()
	// * in front of type means type description. Means here we aren't converting the pointer to actual value but just saying that this is the datatype we expect.
	// * in front of pointer means we want the value corresponding to that address
	// & means give me the address of the variable.
	// turn address into value with *address
	// turn value into address with &value

	// imp another way for by reference. Go will automatically convert to pointer. However not valid for the arguements, only receiver.
	jane.changeFirstNameRef2("Janny")
	jane.printPerson()

	// imp another way for by reference. Go will automatically convert pointer to real object in function. Works both for receiver and arugement
	jane.changeFirstNameRef3("Janny2")
	jane.printPerson()

	// One important thing regarding pointers. Say we pass a slice by value to receiver func then original slice will be changed. The reason is slice is pointer to underlying to array hence changes.

	// When we create a slice Go creates a slice and array data structure for us. Slice has variables ptrToHead, length, capacity. ptrToHead points to the array in the memory. When we call slice we don't refer to slice but array. Go is still pass by value. Hence it creates a new slice for us having 3 new variabes ptrToHead, length, capacity. Values are copied hence ptrToHead is also copied. But ptrToHead still points to the same original array hence no change.

	// Not only slice but other reference types behaves the same way. Slice is reference type variable. Other includes => map, channels, pointers, functions.
	// Value types include => int, float, string, bool, structs. Use pointer to change these things in the function.

	// If you pass a struct by value to a function, a copy of the struct is made, regardless of whether it resides on the heap or stack.

	// Below also works
	// type MyStruct struct {
	// 	value int
	// }

	// func (ms MyStruct) changeNum() {
	// 	ms.value = 5
	// }

	// func createStruct() *MyStruct {
	// 	s := MyStruct{value: 42}
	// 	// s is allocated on the heap because its reference is returned
	// 	return &s
	// }

	// func main() {
	// 	ptr := createStruct()
	// 	// ptr is a pointer to the heap-allocated MyStruct
	// 	fmt.Println(ptr.value) // Output: 42
	// 	fmt.Println((*ptr).value) // Output: 42 see here both above and this lines works

	// ptr.changeNum()
	// fmt.Println(ptr.value)
	// }

}
