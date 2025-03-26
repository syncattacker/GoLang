package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is just a simple function with no args and no return value.")
}


/*
This function accepts two arguments a and b of type integer, and returns the output as integer 
*/
func add(a int, b int) (int) {
	return a + b
}

/*
This function also does addition of two numbers a and b of type integer, adn returns the result in the variable
sum which is of type int.
*/

func addMethod (a int, b int) (sum int) {
	sum =  a + b
	return
}

func main() {
	simpleFunction()
	ans := add(3 , 4)
	fmt.Println("The sum is",ans)
	sum := addMethod(10, 20)
	fmt.Println("Addition by Add Method", sum)
}