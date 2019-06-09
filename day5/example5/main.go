package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"agess"`
	score int
}

func main() {

	stu := Student{Name: "我们", Age: 12, score: 34}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(data))
	}
}
