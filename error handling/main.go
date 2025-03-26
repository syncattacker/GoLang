package main

import "fmt"

func divide(numerator float32, demoninator float32) (float32, error) {
	if demoninator == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return numerator / demoninator, nil
}

func main() {
	fmt.Println("Started error handling in the functions.")
	ans, _ := divide(4, 2)
	fmt.Println(ans)

}