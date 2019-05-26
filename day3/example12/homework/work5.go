package homework

import (
	"fmt"
	"strings"
)

func MaxSum(str string) {

	hasAdd := strings.Index(str, "+")
	if hasAdd != -1 {
		strSlice := strings.Split(str, "+")
		if len(strSlice) == 2 {
			strNumber1 := strings.TrimSpace(strSlice[0])
			strNumber2 := strings.TrimSpace(strSlice[1])
			calc(strNumber1, strNumber2)
		} else {
			fmt.Println("please input a+b")
			return
		}

	}

}

func calc(str1, str2 string) (result string) {

	return
}
