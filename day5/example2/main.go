package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	stu := Student{Name: "nihao "}
	stu.Age = 118
	fmt.Println(&stu.Name)
}
