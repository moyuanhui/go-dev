package main

import (
	"fmt"
	"sort"
)

//排序

func testInt() {
	a := [...]int{3, 4, 5, 10, 1, 4, 2}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testFloat() {
	a := [...]float64{45.3, 23.4, 1.0, 34.65, 76.34}
	sort.Float64s(a[:])
	fmt.Println(a)
}

func testString() {
	a := [...]string{"23d", "fdsf", "hgsdfgsdfgsd", "ds", "gfdgd"}
	sort.Strings(a[:])
	fmt.Println(a)
	index := sort.SearchStrings(a[:], "ffdsf")
	fmt.Println(index)
}

func main() {

	testInt()
	testFloat()
	testString()
}
