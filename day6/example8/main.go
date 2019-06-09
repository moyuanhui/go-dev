package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)
	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()

	stu, ok := iv.(Student)
	if ok {
		fmt.Println("%v %T\n", stu, stu)
	}

}

func testInt(b interface{}) {

	a := reflect.ValueOf(b)
	c := a.Int()
	fmt.Println(c)
}

func testFloat(b interface{}) {

	val := reflect.ValueOf(b)
	val.Elem().SetInt(56)

	c := val.Elem().Int()
	fmt.Println("get value interface{} %d\n", c)
}

func main() {

	stu := Student{"哈哈", 23, 34}
	test(stu)
	var c int = 23
	testFloat(&c)
}
