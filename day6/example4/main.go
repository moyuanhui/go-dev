package main

import "fmt"

func Just(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d param is bool,value is %v", index, v)
		case int, int64, int32:
			fmt.Printf("%d param is int,value is %v", index, v)

		}
	}
}

func main() {

	var b int64 = 23
	Just(b)
}
