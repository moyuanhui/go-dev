package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 互斥锁
var lock sync.Mutex
var rwLock sync.RWMutex

func testMap() {
	var a map[int]int

	a = make(map[int]int, 5)
	a[0] = 190
	a[3] = 5
	a[40] = 13
	a[56] = 546
	a[34] = 34

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[0] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
	time.Sleep(time.Second)
}

func testRwSock() {
	var a map[int]int

	var count int32 //原子操作
	a = make(map[int]int, 5)
	a[0] = 190
	a[3] = 5
	a[40] = 13
	a[56] = 546
	a[34] = 34

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[34] = rand.Intn(100)
			time.Sleep(time.Millisecond * 10)
			lock.Unlock()
		}(a)
	}

	for i := 0; i < 1000; i++ {
		go func(b map[int]int) {
			for {
				rwLock.RLock()
				// fmt.Println(b)
				time.Sleep(time.Millisecond)
				rwLock.RUnlock()
				atomic.AddInt32(&count, 1) //原子操作
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	//testMap()
	testRwSock()
}
