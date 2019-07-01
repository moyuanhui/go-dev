package main

import (
	"fmt"
	"time"
)

// func write(ch chan int) {
// 	for i := 0; i < 1000; i++ {
// 		ch <- i
// 		fmt.Println("put data:", i)
// 	}
// }

// func read(ch chan int) {

// 	for {
// 		b := <-ch
// 		fmt.Println(b)
// 		time.Sleep(time.Second)
// 	}
// }

func calc(taskChan, resultChan chan int) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resultChan <- v
		}
	}
}

func main() {
	chanInt := make(chan int, 100)
	resultChan := make(chan int, 100)
	go func() {
		for i := 0; i < 100000; i++ {
			chanInt <- i
		}
		close(chanInt)

	}()

	for i := 0; i < 8; i++ {
		go calc(chanInt, resultChan)
	}

	go func() {
		for v := range resultChan {
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Second * 10)
}
