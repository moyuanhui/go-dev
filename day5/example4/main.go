package main

import "fmt"

//二叉树
type Student struct {
	Name  string
	Age   int
	Score float32
	Left  *Student
	Right *Student
}

//广度 优先

func trans(root *Student) {

	if root == nil {
		return
	}
	// 前序遍历
	fmt.Println(root)
	trans(root.Left)
	trans(root.Right)

	// // 中序遍历
	// trans(root.Left)
	// fmt.Println(root)
	// trans(root.Right)

	// // 后序遍历
	// trans(root.Left)
	// trans(root.Right)
	// fmt.Println(root)
}

func main() {

	root := Student{Name: "root", Age: 18, Score: 99}
	stu01 := Student{Name: "stu01", Age: 18, Score: 99}

	stu02 := Student{Name: "stu02", Age: 18, Score: 99}

	stu03 := Student{Name: "stu03", Age: 18, Score: 99}
	root.Left = &stu01
	root.Right = &stu02
	stu01.Left = &stu03

	trans(&root)
}
