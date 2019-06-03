package main

import (
	"errors"
	"fmt"
	"time"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// b := 0
	// a := 100 / b
	// fmt.Println(a)
	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}

// append
func main() {
	var a []int
	a = append(a, 10, 20, 383)
	a = append(a, a...)
	fmt.Println(a)
	for {
		test()
		time.Sleep(time.Hour)
	}
}
