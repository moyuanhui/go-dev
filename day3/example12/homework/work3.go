package homework

import (
	"fmt"
	"strings"
)

func IsComeBack(param string) {
	result := ""
	fieldStr := strings.Split(param, "")
	fieldLen := len(fieldStr)
	fmt.Println(fieldLen)
	for i := fieldLen - 1; i >= 0; i-- {
		result += fieldStr[i]
		fmt.Println(fieldStr[i])
	}
	fmt.Println("回文结果是：", result)
	if result == param {

		fmt.Printf("字符串：%s 是回文\n", param)
	} else {
		fmt.Printf("字符串：%s 不是回文\n", param)
	}
}

func Count(str string) bool {

	t := []rune(str)
	fmt.Println(t)
	length := len(t)

	for i, _ := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}
	return true
}
