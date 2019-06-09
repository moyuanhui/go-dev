package main

import (
	"fmt"
	"time"
)

type Car struct {
	name string
	age  int
}

type Train struct {
	Car
	int
	start time.Time
}

func main() {

	var t Train
	t.name = "train"
	t.age = 100
	t.int = 100
	fmt.Println(t)
}
