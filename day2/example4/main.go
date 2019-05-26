package main

import (
	"fmt"
)
func swap(a, b *int){
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func test() {
	a := 100
	fmt.Println(a)
	for i := 0; i < 100; i++ {
		b := i * 2
		fmt.Println(b)

	}
}

func main() {
	a, b := 1, 2
	swap(&a,&b)
	fmt.Printf("a=%d,b=%d \n",a,b)

	test()
}