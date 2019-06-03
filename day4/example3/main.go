package main

import "fmt"

// slice

func test() {
	s1 := new([]int)
	fmt.Println(s1)

	s2 := make([]int, 10)
	fmt.Println(s2)

	*s1 = make([]int, 5)
	(*s1)[0] = 10
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	return
}
func main() {

	test()
}
