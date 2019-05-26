package work

import (
	"fmt"
)

func CalcSx()  {
	for i := 100; i <= 999; i++ {
		one := i % 10
		two := (i / 10) % 10
		three := (i/100) % 10

		var reuslt = one*one*one + two*two*two + three* three * three
		if reuslt == i {
			fmt.Println(reuslt)
		}
		
	}
}