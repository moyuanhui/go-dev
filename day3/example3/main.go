package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := " hello world abc   \n"
	result := strings.Replace(str, "world", "haha", 1)
	fmt.Println(result)

	count := strings.Count(str, "world")
	fmt.Println(count)

	result = strings.Repeat(str, 3)
	fmt.Println("repeat:", result)

	result = strings.ToLower(str)
	fmt.Println("ToLower:", result)

	result = strings.ToUpper(str)
	fmt.Println("ToUpper:", result)

	result = strings.TrimSpace(str)
	fmt.Println("TrimSpace:", result)

	sqlitResult := strings.Fields(str)
	for i := 0; i < len(sqlitResult); i++ {
		fmt.Println(sqlitResult[i])
	}

	splitResult1 := strings.Split(str, " ")
	fmt.Println("split")
	for i := 0; i < len(splitResult1); i++ {
		fmt.Println(splitResult1[i])
	}

	str2 := strings.Join(splitResult1, "123")
	fmt.Println("str2:", str2)

	str2 = strconv.Itoa(1000)
	fmt.Println("Itoa:", str2)

	str3, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("can not convert to int", err)
		return
	} else {
		fmt.Println(str3)
	}
}
