package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int]uint64)
	lock sync.Mutex
)

type task struct {
	n int
}

func test() {
	var i int
	for {
		fmt.Println(i)
		time.Sleep(time.Microsecond)
		i++
	}
}

func calc(t *task) {
	var sum uint64 = 1

	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	//go test()
	// for {
	// 	fmt.Println("i runing in main")
	// 	time.Sleep(time.Microsecond)
	// }
	for i := 0; i < 50; i++ {
		t := &task{n: i}
		go calc(t)
	}
	time.Sleep(time.Second * 5)
	lock.Lock()
	for k, v := range m {
		fmt.Printf("%d=%v\n", k, v)
	}

	lock.Unlock()
}
