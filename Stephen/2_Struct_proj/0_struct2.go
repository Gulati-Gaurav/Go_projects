package main

import "fmt"

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

func Receiver() {
	jane := person3{ // with same name as struct contactInfo
		firstName: "Jane",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "janedoe.com",
			pincode: 110053,
		},
	}

	// using receiver func
	jane.printPerson()

	// IMP receiver is passed by value not reference => will not change original object. Hence will have to use pointers
	// IMP Go is a pass by value language whether you send in arguments or as receiver. See what is C#
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
}
