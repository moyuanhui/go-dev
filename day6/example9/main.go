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

func (p Student) Print() {
	fmt.Println("ahh 323")
}

func (p Student) SET(name string, age int, score float32) {
	fmt.Println("ahh ")
}

func TestStruct(p interface{}) {

	val := reflect.ValueOf(p)
	kind := val.Kind()
	if kind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	fieldNum := val.NumField()
	fmt.Printf("struct has %d field \n", fieldNum)

	for i := 0; i < fieldNum; i++ {
		fmt.Printf("%d %v\n", i, val.Field(i))
	}

	methodNum := val.NumMethod()
	fmt.Printf("struct has %d method \n", methodNum)

	var params []reflect.Value
	val.Method(0).Call(params)
}

func main() {

	stu := Student{"你好", 23, 34}
	TestStruct(stu)
	fmt.Println(stu)

}
