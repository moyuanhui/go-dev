package main

import (
	"go_dev/day5/sortexample"
)

func main() {
	arr := [...]int{34, 23, 5, 66, 22, 77, 2, 7, 1}
	arr1 := [...]int{34, 23, 5, 66, 22, 77, 2, 7, 1}

	arr2 := [...]int{34, 23, 5, 66, 22, 77, 2, 7, 1}

	arr3 := [...]int{34, 23, 5, 66, 22, 77, 2, 7, 1}

	sortexample.TestMaopao(arr[:])
	sortexample.TestQuick(arr1[:])
	sortexample.TestSelect(arr2[:])
	sortexample.TestInsert(arr3[:])
}
