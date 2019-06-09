package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string `json::"name"`
	Id   int
	Age  int
}

type StudentArray []Student

func (s StudentArray) Len() int {
	return len(s)
}

func (s StudentArray) Less(i, j int) bool {
	return s[i].Name > s[j].Name
}

func (s StudentArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	var stus StudentArray
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", rand.Intn(10)),
			Id:   rand.Intn(100),
			Age:  rand.Intn(100),
		}
		stus = append(stus, stu)
	}
	for _, v := range stus {
		fmt.Println(v)
	}

	fmt.Println("排序结果：")
	sort.Sort(stus)

	for _, v := range stus {
		fmt.Println(v)
	}
}
