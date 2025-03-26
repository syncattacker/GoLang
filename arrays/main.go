package main

import "fmt"

func main() {
	fmt.Println("Array Study")
	var names [5]string
	names[0] = "Prince"
	names[1] = "John"
	fmt.Println(names)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Length of Numbers Array Is", len(numbers))
	fmt.Println(numbers[2])
}