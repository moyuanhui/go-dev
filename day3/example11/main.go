package main

import (
	"fmt"
)

func add(a int, arg ...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

func main() {

	sum := add(10, 20, 30)
	fmt.Println("test")
	a := 10
	defer fmt.Println(a)

	a = 0
	fmt.Println(a)
	fmt.Println(sum)
}
