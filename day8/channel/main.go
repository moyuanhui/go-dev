package main

import "fmt"

type Student struct {
	name string
}

func main() {
	// var inchan chan int = make(chan int, 10)
	// inchan <- 10
	// var mapChan chan map[string]string
	// mapChan = make(chan map[string]string, 10)
	// m := make(map[string]string, 16)
	// m["stu01"] = "001"
	// m["stu02"] = "002"
	// mapChan <- m

	// var chanStu chan stu

	var chanStu chan interface{}
	chanStu = make(chan interface{}, 10)
	stu := Student{name: "stuff1"}
	chanStu <- &stu

	stu01 := <-chanStu
	// fmt.Println(Student(stu01))
	stu02, ok := stu01.(*Student)
	if !ok {
		fmt.Println("can not ")
		return
	}
	fmt.Println(stu02)
}
