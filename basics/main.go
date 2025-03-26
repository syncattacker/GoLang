package main

import "fmt"

func main() {
	// Print statement in Go!
	fmt.Println("Hello from Go!")

	// Variables in Go!
	// Variables with const keyboard cannot be reassigned with any values
	const name string = "John Doe"
	// name = "New Name", This won't work as name is a constant type variable.
	const age int = 21
	const weight float32 = 64.965596 // Accepts upto 6 decimal places, then rounds off
	const height float64 = 65.8986569666 // Accepts upto 15 decimal places
	const isDecided bool = true
	fmt.Println(name, age, weight, height, isDecided)

	// Reassignment of Variables in Go!
	var myName string = "James"
	myName = "Rahul"
	fmt.Println(myName)

	/* Varaibles with defined types are strictly coupled while compile time, making it hard to change the data type 
	ex : const age int = 30 (age cannot have any other values `const myAge int = "30"`)
	For this go have other kind of variable declaration.
	*/ 

	var myAge = 30
	fmt.Printf("Type of myAge is %T\n", myAge)
	var otherAge = "30"
	fmt.Printf("Type of myAge is %T\n", otherAge)

	// Other kind of variable assignment supported by Go!

	newAge := "30"

	fmt.Printf("Type of myAge is %T", newAge)
	

	  
	

}