package main

import "fmt"

// to not create a struct for whole package but just inside a function
func temp_structs_go() {
	obj := struct {
		name  string
		price float64
	}{
		name: "Gaurav",
		price: 12,
	}

	fmt.Println(obj.name)
	fmt.Println(obj.price)
}
