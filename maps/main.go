// package main

// import "fmt"

// func main() {

// 	colors := map[string]string {
// 		"red": "#123",
// 		"blue": "#213",
// 		"green": "#431",
// 	}

// 	colors["yellow"] = "#432"
// 	printMap(colors)
// }

// func printMap(m map[string]string) {
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}
// }
package main

import (
	"fmt"
)

func Index[T comparable](s []T, x T) int {
	for i, val := range s {
		if val == x {
			return i
		}
	}
	return -1
}

func main() {
	s := []int{3, 4, 5, 6, 7}
	fmt.Println(Index(s, 5))
}