package main

import "fmt"

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

func Initialisation() {
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
	// It is important to put comma in instance initialisation not in struct declaration above
	// Also commas need to placed after all properties including last property

	// 3rd instant initialisation.
	// If we don't put values, struct puts 0 value by default. means empty string for string, 0 for int, float. false for bool
	var bill person
	// to put / change values
	bill.firstName = "Bill"
	bill.lastName = "Gates"

	// var billy person2
	// Can't do below since, Only pointers, slices, maps, channels, and interfaces can be compared to nil.
	// fmt.Println(billy.contact==nil)
	// If you do not explicitly initialize the contact field in person2, it will be assigned its zero value by default. 
	// In Go, all struct fields are automatically initialized to their respective zero values. Even on multiple nesting of structs no problem of nil (unless the data members are the one mentioned above)

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
	
	fmt.Printf("%+v \n", john)
}
