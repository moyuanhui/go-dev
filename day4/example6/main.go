package main

import "fmt"

func fab(n int) {
	var a []int64
	a = make([]int64, n)

	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for _, v := range a {
		fmt.Println(v)
	}
}

func testMultiArrry() {
	a := [...][3]int{{1, 2, 3}}
	fmt.Println(a)
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}

func main() {
	fab(10)

	var age0 [5]int = [5]int{1, 2, 3}
	var age1 = [5]int{1, 2, 3}
	var age2 = [...]int{1, 2, 3, 4}
	var str = [5]string{3: "hello", "4:world"}
	a := [...]int{4: 4, 5: 3}

	fmt.Println(age0)
	fmt.Println(age1)

	fmt.Println(age2)

	fmt.Println(str)
	fmt.Println(a)
	testMultiArrry()

}
