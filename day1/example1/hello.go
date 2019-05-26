package main 

import (
	"fmt"
	"time"
)
func add(a,b int) int {
	c := a + b
	fmt.Println(c)
	return c
}

func main()  {
	fmt.Printf("hello world\n")
	// for index := 0; index < 100; index++ {
	// 	go test_goroute(index,300)
	// }
	// // for index := 0; index < 100; index++ {
	// // 	go test_print(index)
	// // }	

	fmt.Println("start")
	go test_pipe()
	fmt.Println("end")
	time.Sleep(100 * time.Second)

	add(1,2)
}