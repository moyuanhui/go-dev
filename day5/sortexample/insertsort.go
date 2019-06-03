package sortexample

import "fmt"

func TestInsert(array []int) {

	for i := 0; i < len(array); i++ {
		j := i
		target := array[i]
		for target < array[j-1] || j > 0 {
			array[j] = array[j-1]
			j--
		}
		array[j] = target
	}
	fmt.Println("插入排序结果")
	fmt.Println(array)
}
