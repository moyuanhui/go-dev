package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}) {

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
}
func main() {

	a := 32
	test(a)
}
