package main

import "fmt"

// The syntax to declare a type in go. Can think of as deck type extends string. deck now has all the properties of string class and we can further include more properties ourselves
type cars string

// c cars is receiver (means the object calling the function). Means any variable of type cars gets access to the basics_carName() func. By convention receiver is referred with 1-2 letter abreviation.
// receiver functions are the reason we created separate type. These help to associate a function with a type
// type of the receiver can't be built-in datatypes.
// c is not actual instance but copy of it
// can think of c as 'this' of c++.
func (c cars) basics_receiver_func() {
	fmt.Println(c)
}

// to return multiple things
func deal(c cars) (string, cars) {
	return string(c), c // this is how we convert one type to another or type conversion in go
}

// TypeWeWant(TypeWeHave)
