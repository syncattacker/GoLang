package main

import "fmt"

func main() {
	day := 30
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winter Season")
	case "April", "May", "June":
		fmt.Println("Spring Season")
	default:
		fmt.Println("Other Season")
	}

	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 25:
		fmt.Println("Hot")
	default:
		fmt.Println("Unknown Temperature")
		
	}
}