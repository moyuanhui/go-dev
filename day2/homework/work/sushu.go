package work

import (
	"fmt"
)

//计算素数
func CalcSuShu()  {
	for n := 101; n <= 200; n++ {
		flag := true
		for i := 2; i <= n-1; i++ {
			if n%i == 0 {
				flag = false
				break
			}
		}
		if flag == true {
			fmt.Printf("n[%d] is 素数\n", n)
		} 
		// else {
		// 	fmt.Printf("n[%d] is not 素数\n", n)
		// }
	}
}