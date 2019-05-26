package main

import (
	"fmt"
)

func main()  {
	var a int = 100
	var b bool
	var c byte = 'a'
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)

	str := fmt.Sprintf("%d", a)
	fmt.Printf("%q",str)
}