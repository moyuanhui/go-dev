package main

import (
	"fmt"
	"time"
)

const (
	Female = 2
	Man = 1
)

func main() {
	for {
		second := time.Now().Unix()
		if second % Female == 0 {
			fmt.Println("Female")
		} else {
			fmt.Println("man")
		}
		time.Sleep(10 * time.Millisecond)
	}

}
