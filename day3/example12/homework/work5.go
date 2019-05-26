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
			result := calc(strNumber1, strNumber2)
			fmt.Println(result)
		} else {
			fmt.Println("please input a+b")
			return
		}

	}

}

func calc(str1, str2 string) (result string) {

	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}
	index1 := len(str1) - 1
	index2 := len(str2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {

		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'
		sum := int(c1) + int(c2) + left

		if left >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c1 := str2[index2] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}
	if left == 1 {
		result = fmt.Sprintf("%c%s", left, result)
	}

	return
}
