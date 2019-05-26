package main

import(
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main()  {
	for i :=0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	for i := 0; i < 100; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i :=0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}

}