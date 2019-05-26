package main

import (
	"fmt"
)

// 反转字符串
func reverse(str string) string {
	var result string 
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

// append实现反转
func reverse1(str string) string {
	tmp := []byte(str)
	var result []byte
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = append(result, tmp[strLen-i-1])
	}
	return string(result)
}

func main()  {
	str1 := "hello"
	str2 := "world"

	str3 := str1 + " " + str2
	str4 := fmt.Sprintf("%s %s",str1, str2)
	fmt.Println(str3)
	fmt.Println(str4)

	fmt.Println("len(str3)=",len(str3))

	substr := str3[4:]
	fmt.Println(substr)

	result := reverse(str3)
	fmt.Println(result)

	result1 := reverse1(result)
	fmt.Println(result1)
}