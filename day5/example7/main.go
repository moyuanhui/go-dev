package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (p *Student) init(name string, age, scorce int) {

	p.Name = name
	p.Age = age
	p.Score = scorce
	fmt.Println(p)
}

func main() {
	var stu Student
	stu.init("nihao", 1, 2)
	fmt.Println(stu)
}
