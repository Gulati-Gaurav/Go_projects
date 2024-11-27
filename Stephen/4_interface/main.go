package main

import "fmt"

// type englishBot struct{}
// type spanishBot struct{}

// func (eb englishBot) getGreeting() string {
// imp since eb is not used we can omit it like below but can't remove brackets
// func (englishBot) getGreeting() string{
// 	return "Hello World"
// }
// func (sb spanishBot) getGreeting() string{
// 	return "Hola"
// }

// imp unlike in other languages, go doesn't support multiple functions with same name unless different receivers.
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }

// => Now we will use interfaces to prevent the code duplicacy of the printGreeting which has similar logic.
type bot interface {
	getGreeting() string
}

// To whom it may concern...
// Our program has a new type called 'bot'
// If you are a type in this program with a function called 'getGreeting' and it returns a string then you are now an honorary member of type 'bot'. You are a type of 'bot'. You have type 'bot'
// Now that you're also an honorary member of type 'bot', you can now call this function called 'printGreeting'
// All the methods of interface (eg. getGreeting in this case) are receiver type methods (either value receiver or pointer receiver) (ofc since if not then how do you associate yourself with that function).
// Further the functions (eg. printGreeting) can't be receiver type since interfaces cannot have methods with implementations.

// We can also mention list of the parameters of the function desired eg.
// Can also mention different functions eg.
// Can have other interface arguement or return type
// An interface can implement another interface eg.
type bot2 interface {
	name(string, int) (map[string]string, int)
	printRoll() string
	printRoll2(bot) bot
	bot
}

// There are 2 types of 'type' in GO. Concrete type and Interface Type.
// Concrete type includes map, struct, int etc. Means we can create a value out of it directly and change it and create new copies of it.
// Can't create a value (a variable of interface type like bot)

// A type can implement multiple interfaces.

// The term "generic type" refers to a type that can work with multiple types of data, abstracting over the specific data types it operates on.
// Interfaces are not replacement for generic types
// Now have been added in go. eg. in c++
// template<typename T>
// void swap(T& a, T& b)

// Interfaces are implicit i.e. we don't need to implicitly mention that a struct/type is implementing a interface. Boon as well as bane as we don't get to know which interface a struct is implementing.

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello"
}
func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	// eb := englishBot{}
	// sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)
}
