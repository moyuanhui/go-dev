package sortexample

import "fmt"

func TestMaopao(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println("冒泡排序结果")
	fmt.Println(array)
}
