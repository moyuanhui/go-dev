package sortexample

import "fmt"

func TestSelect(array []int) {

	minIndex := 0
	for i := 0; i < len(array); i++ {
		minIndex = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
	fmt.Println("选择排序结果")
	fmt.Println(array)
}
