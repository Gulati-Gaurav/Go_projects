package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// AnotherWay()
	// AnotherWay1()
	// AnotherWay1_5()
	// AnotherWay2()
	AnotherWay3()
}

// 2 ways to make http request
// 1. Use the net package and in it http package
// 2. Use fibre (in room mapping)
func AnotherWay() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Found error", err)
		os.Exit(1)
	}
	// This contains status, StatusCode etc
	// body is of type interface called ReadCloser. This interface further implements Reader and Closer interface which requires reader and closer functions to fulfill.
	// IMP when a variables' type is an interface and you wanna use it. You can only use it via the functions it implements

	// Reader func accepts byte array and fills it with the response. Hence not only http request but we can also have many other struct implement these Reader interface which gives us a common way of interaction. See Image.
	// Reader solves the problem of same logic different codes due to just type of arguement by making the type of arguement common to []byte

	// Interface can not have data members i.e. can not mention a shape interface to have side as a data member. Workaround is to make a function which returns side.

	httpBody := make([]byte, 22792) // reason is Read doesn't expand slice's size
	n, err2 := resp.Body.Read(httpBody)
	if err2 != nil {
		fmt.Println(err2) // gives EOF = End of File . So either ignore because go is sending after encoding in gzip or see AnotherWay2
		// Will not give EOF if you make the byte slice of exactly the size of the body
		fmt.Println(n)
	}
	fmt.Println(string(httpBody))
	resp.Body.Close()

	// Why close the body ?
	// When you get the data (response body) from a website, you need to "close" it to free up resources and avoid wasting them.
}
