package main

import (
	"fmt"
)

func main() {

	a := 10
	fmt.Println(&a)

	// var p *int
	p := &a
	fmt.Println(*p)

	*p = 15
	fmt.Println(*p)
}
