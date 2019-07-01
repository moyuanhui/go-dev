package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(j)
		}(i)
	}
	time.Sleep(time.Second)
}
