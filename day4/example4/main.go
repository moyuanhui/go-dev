package main

import (
	"fmt"
	"time"
)

func recusive() {
	fmt.Println("调用了")
	time.Sleep(time.Second)
	recusive()
}

func main() {

	recusive()
}
