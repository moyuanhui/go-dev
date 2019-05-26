package main

import (
	"fmt"
	"strings"
)

// 字符串
func main()  {
	strHttp := "http://www.baidu.com"

	//判断是以字符串开头 HasPrefix
	isPrefixHttp :=  strings.HasPrefix(strHttp,"http")
	if isPrefixHttp == true {
		fmt.Printf("strHttp is prefix http\n")
	} else {
		fmt.Printf("strHttp is not prefix http\n")
	}

	// 字符串以什么结尾 HasSuffix
	isSuffix := strings.HasSuffix(strHttp, "com")
	if isSuffix == true {
		fmt.Printf("strHttp is Suffix com\n")
	} else {
		fmt.Printf("strHttp is not Suffix com\n")
	}

	// 第一次出现Index
	index := strings.Index(strHttp, "baidu")
	fmt.Println("baidu第一次出现位置:",index)
	
	//最后出现位置 LastIndex
	lastIndex := strings.LastIndex(strHttp, "baiduf")
	fmt.Println("baidu最后一次出现位置:",lastIndex)

	//Replace 字符串替换
	replaceResult := strings.Replace(strHttp,"http","haha",-1)
	fmt.Println(replaceResult)

	//去掉空格TrimSpace
0	trimSpace := strings.TrimSpace(" hahah ")
	fmt.Printf("%#v",trimSpace)

}