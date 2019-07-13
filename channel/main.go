package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool,2)
	go func(){
		fmt.Println("go test01")
		time.Sleep(2 * time.Second)
		fmt.Println("channel finish")
		ch <- true
	}()
	<-ch

	fmt.Println("go test02")
	// time.Sleep(2*time.Second)
	fmt.Println("main finish")
}
