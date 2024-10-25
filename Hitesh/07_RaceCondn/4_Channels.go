package main

import "fmt"

func channels() {
	ch := make(chan int)
	// ch2 := make(chan int, 2) // to make channel with buffer in cases when you don't want number of listner = number of speaker

	// arrows are always left sided in go

	ch <- 12
	fmt.Println(<-ch)
	// this will make the program to keep waiting forever since channel only allows to put value when someone is listening to it
	// interchanging also brings the same deadlock condition

	// close(ch) // to stop accepting anything in the channel

}
