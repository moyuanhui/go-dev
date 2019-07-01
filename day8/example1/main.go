package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 1000; i++ {
		ch <- i
		fmt.Println("put data:", i)
	}
}

func read(ch chan int) {

	for {
		b := <-ch
		fmt.Println(b)
		time.Sleep(time.Second)
	}
}

func main() {
	chanInt := make(chan int, 10)
	go write(chanInt)
	go read(chanInt)

	time.Sleep(time.Second * 100)
}
