package main

import (
	"fmt"
	"math/rand"
)

// 单链表
type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func trans(p *Student) {

	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func testInsertTail(head *Student) {
	tail := head
	for i := 0; i < 10; i++ {
		stu := Student{Name: "hua2", Age: rand.Intn(100), Score: rand.Float32() * 100}
		tail.next = &stu
		tail = &stu
	}
	trans(head)
}

func testInsertHead(head *Student) {

	for i := 0; i < 10; i++ {
		stu := Student{Name: "hua2", Age: rand.Intn(100), Score: rand.Float32() * 100}
		stu.next = head
		head = &stu
		// head.next = stu.next
	}
	trans(head)
}
func main() {

	head := Student{Name: "hua1", Age: 18, Score: 100}
	// stu1 := Student{Name: "hua", Age: 18, Score: 99}
	// stu2 := Student{Name: "hua", Age: 18, Score: 88}
	// stu3 := Student{Name: "hua", Age: 18, Score: 77}

	// head.next = &stu1

	// p := &head
	// stu1.next = &stu2
	// stu2.next = &stu3

	// trans(p)

	// testInsertTail(&head)
	testInsertHead(&head)

}
