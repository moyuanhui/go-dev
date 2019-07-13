package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s", p.path)
}

func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       "test.sls",
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}

func main() {
	err := Open("C:/ss.txt")
	if err != nil {
		fmt.Println(err)
	}
	v, ok := err.(*PathError)
	if ok {
		fmt.Println("get path error", v)
	}
}
