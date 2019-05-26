package homework

import (
	"fmt"
)

//1000以内的完数
func CalcWanShu(maxCount int) {
	fmt.Println("完数有一下内容：")
	for i := 1; i <= maxCount; i++ {
		sum := 0
		avg := i / 2
		for j := 1; j <= avg; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum == i {
			fmt.Println(i)
		}
		sum = 0
	}
}
