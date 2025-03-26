package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6,7,8,9,10,11,12,13,14)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Printf("Type of Numbers Is : %T\n", numbers)

	// Slice through make function
	myNumbers := make([]int, 3, 5 )
	myNumbers = append(myNumbers, 1, 2, 3)
	fmt.Println("Slice", myNumbers)
	fmt.Println("Length", len(myNumbers))
	fmt.Println("Capacity", cap(myNumbers))
}