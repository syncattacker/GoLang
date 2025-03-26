package main

import "fmt"

func main() {
	x := 10
	if x > 0 {
		fmt.Println("The value of X is greater than 0.")
	} else {
		fmt.Println("The value of X is less than 0.")
	}

	z := 10
	if z > 10 {
		fmt.Println("Z is greater than 10.")
	} else if z > 5 {
		fmt.Println("Z is greater than 5.")
	} else {
		fmt.Println("Z is smaller than 5.")
	}
}