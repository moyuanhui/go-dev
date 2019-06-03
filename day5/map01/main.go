package main

import "fmt"

func testMap() {

	a := make(map[string]string, 10)
	a["abc"] = "efg"
	fmt.Println(a)

	b := map[string]int{"aabb": 23}
	fmt.Println(b)
}

func testMap3() {

	a := make(map[string]map[string]int, 8)
	b := map[string]int{"dd": 23}
	a["key1"] = b
	a["key2"] = make(map[string]int, 4)
	a["key2"]["key3"] = 23
	fmt.Println(a)

	val, ok := a["key1"]
	if ok {
		fmt.Println("存在key1", val)
	} else {
		a = make(map[string]map[string]int)
	}

	for key, val := range a {
		fmt.Println(key, val)
	}
}

func testMap5() {
	var a []map[int]int

	a = make([]map[int]int, 5)
	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 10
	fmt.Println(a)
}

func main() {

	testMap()
	// testMap2()
	testMap3()
	testMap5()
}
