package main

import "fmt"

// 闭包
func Adder() func(int) int {

	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func main() {

	var f = Adder()
	fmt.Print(f(1), "-")
	fmt.Print(f(20), "-")
	fmt.Print(f(300))
}
