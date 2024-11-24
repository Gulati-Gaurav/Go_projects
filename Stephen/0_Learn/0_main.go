package main

import "fmt"

func main() {

	fmt.Println("Hi There")         // To Print
	fmt.Println("Hi There" + "adf") // To Print
	fmt.Println("Hi There", 2)      // To Print multiple (spaces get added in b/w)

	variables()
	arrays_slices_for_discard_subset()

	// Type and their functions
	car := cars("Maruti") // or var car cars = "Maruti"
	car.basics_receiver_func()

	car1, _ := deal(car)
	fmt.Println(car1)

	// One nice thing about go is that if just over a function you mention a comment, you will be able to see that message when you hover over that function anywhere (just like summary in c#)
}
