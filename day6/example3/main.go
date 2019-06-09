package main

import "fmt"

// 接口类型,类型断言
func Test(a interface{}) {

	b, ok := a.(int)
	if ok {
		b += 3
		fmt.Println(b)
	} else {
		fmt.Println("false")
	}

}

func main() {
	b := "ff12"
	Test(b)
}
