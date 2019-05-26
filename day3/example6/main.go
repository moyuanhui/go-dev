package main

import "fmt"

func PrintResult(n int) {

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("test")
	PrintResult(6)
}
