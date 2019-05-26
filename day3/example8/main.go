package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	c := add
	fmt.Println(c)

	sum := c(1, 2)
	fmt.Println(sum)

	// if c == add {
	// 	fmt.Println("c equal add")
	// }

}
