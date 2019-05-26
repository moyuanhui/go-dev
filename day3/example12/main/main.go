package main

import (
	"bufio"
	"fmt"
	"go_dev/day3/example12/homework"
	"os"
)

func main() {

	fmt.Println("九九乘法表")
	homework.PrintNine()
	fmt.Println("1000以内的完数")
	homework.CalcWanShu(1000)
	fmt.Println("判断回文")
	homework.IsComeBack("112211  ")
	str := "上海自来水来自海上  "
	result := homework.Count(str)
	if result == true {
		fmt.Printf("%s 是回文", str)
	} else {
		fmt.Printf("%s 不是回文", str)
	}

	render := bufio.NewReader(os.Stdin)
	_, _, err := render.ReadLine()
	if err != nil {
		fmt.Println("read from console err", err)
	}
	workCount, spaceCont, _, _ := homework.CountField("fadfadf发电房  gg!  ")
	fmt.Println("workCount = ", workCount)
	fmt.Println("spaceCont = ", spaceCont)

}
