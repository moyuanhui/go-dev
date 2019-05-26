package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Microsecond * 100)
}

func main() {

	//格式化时间
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:03"))
	// msg := now.Format("2017-20-32 12:23:23")

	// 统计时间耗时
	startTime := time.Now().UnixNano()
	test()
	endTime := time.Now().UnixNano()
	fmt.Printf("cost:%d us\n", (endTime-startTime)/1000)

}
