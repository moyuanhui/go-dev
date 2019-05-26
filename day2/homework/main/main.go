package main  

import (
	"fmt"
	"go_dev/day2/homework/work"
)

func main()  {
	fmt.Println("第二天课后作业")

	fmt.Println("第一道题：打印出101-200之间有多少个素数")
	work.CalcSuShu()

	fmt.Println("第二道题：打印出100-999之间的水仙花数")
	work.CalcSx()


	fmt.Println("第三道题：求对于一个n的阶乘之和")
	var n int
	// fmt.Scanf("%d", &n)
	n = 5
	work.CalcJieCheng(n)

}