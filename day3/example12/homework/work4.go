package homework

import "fmt"

// 判断字符
func CountField(str string) (worldCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		fmt.Printf("%v \n", v)
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			worldCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}
