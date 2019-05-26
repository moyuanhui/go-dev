package main

import "fmt"

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op add_func, a, b int) int {
	return op(a, b)
}

func main() {

	c := add
	fmt.Println(c)
	sum := operator(c, 100, 200)
	fmt.Println(sum)
}
