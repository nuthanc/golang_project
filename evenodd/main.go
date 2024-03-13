package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		evenOrOdd := "even"
		if number % 2 == 1 {
			evenOrOdd = "odd"
		}
		fmt.Printf("%v is %v\n", number, evenOrOdd)
	}
}

