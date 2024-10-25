package main

import (
	"fmt"
	// "io"
	"os"
)

func functions() {
	// Can't have named functions inside functions

	// Anonymous functions (lamda functions) are like -
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	// Anonymous function can be passed as argument
	numbers := []int{1, 2, 3, 4, 5}
	result := filter(numbers, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(result)

	// IIFEs are here too.
	func() {
		fmt.Println("Hello")
	}()
	greet("Gaurav")

	// Difference between functions and methods = functions inside classes are called methods
}
func filter(slice []int, f func(int) bool) []int {
	var result []int
	for _, num := range slice {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}

func deferGo() {
	// defer invokes a function whose execution is deferred to the moment the surrounding function returns (either end of body or return or panic) and is not invoked immeditealy where written

	// If 2 deferred functions call in the same functions, the order will be LIFO.

	// Imagine like it cuts the line from there and pastes in the end
}

func filesInteraction() {
	// creation, open utility lies in os hands
	// i/o operations are in the hand of io
	// however now go has given os powers to make it easy

	content := "This needs to go into file"
	path := "./myfile.txt"

	// file, err := os.Create("./myfile.txt")
	// checkErrorNil(err)

	// len, err := io.WriteString(file, content)
	// checkErrorNil(err)
	// fmt.Println("length = ", len)

	// defer file.Close()

	// or

	err := os.WriteFile(path, []byte(content), 0666) // closing the file is handled inside this
	checkErrorNil(err)

	readFile("./myfile.txt")
}

func readFile(path string) {
	// file, err := os.Open(path)
	// checkErrorNil(err)

	// data, err := io.ReadAll(file)
	// checkErrorNil(err)

	// fmt.Println(string(data))

	//or

	data, err := os.ReadFile(path)
	checkErrorNil(err)
	fmt.Println(string(data))
}

type Shape interface{ Area() int }
type Rectangle struct{}
type Circle struct{}

func (Rectangle) Area() int { return 1 }
func (Circle) Area() int    { return 1 }

func typeAssertion() {
	// is a mechanism to check the dynamic type of a value during runtime and extract its underlying concrete value.
	// It's primarily used with interface values, where the actual type of the underlying data is unknown.

	var i interface{} = "hello"
	s, ok := i.(string)
	if ok {
		fmt.Println("i is a string:", s)
	} else {
		fmt.Println("i is not a string")
	}

	DoSomething := func(s Shape) {
		switch v := s.(type) {
		case Rectangle:
			fmt.Println("Rectangle area:", v.Area())
		case Circle:
			fmt.Println("Circle area:", v.Area())
		default:
			fmt.Println("Unknown shape")
		}
	}
	DoSomething(Shape(Rectangle{}))

	// you don't need to add break after every case in a Go switch statement.

	// Unlike many other programming languages, Go implicitly breaks out of a switch case after the code block is executed. This means that once a matching case is found, the subsequent cases are ignored.
}

// this is a common way to handle the errors in go
func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
