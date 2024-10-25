package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Print() {
	// To print without enter
	fmt.Print("sdf") // instead of Println
	// in fact Println is nothing but inserts \n at the end of string
}

func DeclarePrintMultipleVariables() {
	var num, i, j int
	fmt.Println(num, i, j)
}

func TakeInputFromUser() {
	// How to take input from user
	// Scan - receives the user input in raw format as space-separated values and stores them in the variables. Newlines count as spaces.
	var num, i, j int
	fmt.Scan(&num)
	fmt.Println(num)

	fmt.Scan(&i, &j) // to take multiple inputs
	fmt.Println(i, j)

	// Scanln() function is similar to Scan(), but it stops scanning for inputs at a newline (at the press of the Enter key.) -> will put the default value of int in j in case pressed enter after i.
	fmt.Scanln(&i, &j)
	fmt.Println(i, j)

	// Scanf() function receives the inputs and stores them based on the determined formats for its arguments.
	fmt.Scanf("%v %v", &i, &j)
	// In this example, the inputs are received in exactly the same way defined in Scanf() formatting.
	// This means that the code expects inputs to be received with one space between them. Like: 1 2.

	fmt.Scanf("%v\n%v", &i, &j)
	// This example receives the values of i and j in separate lines and then prints them:
}

func TakeInputFromUserBufio() {
	// bufio is stores input and output in a buffer as name says
	// In the arguement of NewReader you mention what the source of the reader will be.
	// just delaration. not takes input here
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating ")

	// below is called comma error or error ok syntax
	resp, err := reader.ReadString('\n') // \n is for how long do you wanna read. Means read till new line
	if err != nil {
		fmt.Println("error = ", err)
		os.Exit(1)
	} else {
		fmt.Println("Thanks for rating ", resp)
	}
}

func Conversion() {
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		inp = strings.TrimSpace(inp) // reason being \n also becomes a part of the string.
		rating, err := strconv.ParseInt(inp, 10, 32)
		if err != nil {
			fmt.Println(err)
		} else {
			rating++
			fmt.Println(rating)
		}
	}
}

func Time() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Nanosecond())
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // This is unique to go, this is the format we have to use to mention the format we want
	createdDate := time.Date(2024, time.September, 18, 5, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
}

func MemoryMng() {

	// 2 types of memory allocation - new() and make()
	// In both you get a memory address
	// new() => allocates new memory but no initialisation
	// make() => allocates new memory and initialises

	// new() => zeroed storage (can't put data initally)
	// make() => non-zeroed storage

	// Garbage collection happens automatically
	// when a variable goes out of scope or is nil it qualifies for memory management

	// runtime package is available by go to see info like memory available etc.

	fmt.Println("hello")
}

func Slices() {
	numbers := []int{1, 3, 4, 5}
	numbers = append(numbers[1:3])

	highScores := make([]int, 4) // to make slices of fixed slices
	fmt.Println(highScores)

	sort.Ints(highScores) // to sort

	// How to remove from an array based on index
	numbers2 := []string{"JavaScript", "GoLang", "Ruby", "C#"}
	numbers2 = append(numbers2[:2], numbers2[2+1:]...)

	// ... operator in Go (Golang) is used for variadic functions. A variadic function can accept a variable number of arguments.

	// You can define a function to accept a variable number of arguments by using the ... operator before the type of the last parameter. This parameter will then be a slice of the specified type.

	// func sum(nums ...int) int {
	// 	total := 0
	// 	for _, num := range nums {
	// 		total += num
	// 	}
	// 	return total
	// }

	// Calling a Variadic Function:
	// You can call a variadic function with any number of arguments.
	// total := sum(1, 2, 3, 4)
	// fmt.Println(total) // Output: 10

	// Now function as a spread operator

	// Passing a Slice to a Variadic Function:
	// If you already have a slice and want to pass it to a variadic function, you can do so by appending ... after the slice.
	// nums := []int{1, 2, 3, 4}
	// total := sum(nums...)

	// also used for the substring
	str := []int{1, 2, 3, 4, 5}
	str = append(str[:2], str[3:]...)
}

func IfElse() {
	// when you wanna declare and check both
	if num := 10; num%2 == 0 {
		fmt.Println("Even Number")
	}

	num := 1
	switch num {
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("Pata nhi")
	}
	// can put fallthrough - if one case satisfies the below case will also run even if doesn't match the condition

	// In Go, the increment (++) and decrement (--) operators are only supported in the postfix form. This means that i++ and i-- are valid, but ++i and --i are not.

	// break and continue and goto also in go

	// while loop similar
	for num < 10 {

	}

	// goto also there but don't use
}
