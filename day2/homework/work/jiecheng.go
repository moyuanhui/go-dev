package work

import (
	"fmt"
)

func CalcJieCheng(m int)  {
	result := 0
	s := 1
	for i := 1; i <= m; i++ {
		s = s * i
		fmt.Printf("%d != %v\n",i,s)
		result += s
	}
	
	fmt.Println("n的阶乘=",result)
}