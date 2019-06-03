package main

import (
	"fmt"
	"sort"
)

func testMap5() {
	var a map[int]int

	a = make(map[int]int, 5)
	// if a[0] == nil {
	// 	a[0] = make(map[int]int)
	// }
	// a[0][10] = 10
	// fmt.Println(a)
	a[0] = 190
	a[3] = 5
	a[40] = 13
	a[56] = 546
	a[34] = 34
	fmt.Println(a)

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys[:])

	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func main() {

	testMap5()
}
