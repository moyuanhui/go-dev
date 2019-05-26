package main

import (
	"fmt"
	"strconv"
)

// 水仙花数
func main()  {
	var str string
	fmt.Scanf("%s", &str)

	result := 0
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		num := int(str[i] - '0')
		result += num * num * num
	}

	number,err := strconv.Atoi(str)

	if err != nil {
		fmt.Printf("can not convert %s to int\n", str)
		return
	}

	if result == number {
		fmt.Printf("%d is shuixianhuashu\n", number)
	} else {
		fmt.Printf("%d is not shuixianhuashu\n", number)
	}
}