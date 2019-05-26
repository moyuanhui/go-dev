package main

import "fmt"

func calc(a, b int) (sum, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

func main() {

	sum, avg := calc(10, 20)
	fmt.Printf("sum = %d,avg = %d", sum, avg)
}
