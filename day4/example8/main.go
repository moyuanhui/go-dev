package main

import "fmt"

func testSlice() {
	var arr = [...]int{1, 2, 3, 4, 5}
	var slice []int
	slice = arr[2:]
	fmt.Println(slice)
}

func main() {

	testSlice()
}
