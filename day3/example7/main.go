package main

import (
	"fmt"
)

func main() {

	str := "hello world, 椎间盘"
	for index, val := range str {
		fmt.Printf("index[%d] val[%c] len[%d],%s\n", index, val, len([]byte(string(val))), string(val))
	}
}
