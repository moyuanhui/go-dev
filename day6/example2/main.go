package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("read data")
}

func (f *File) Write() {
	fmt.Println("write data")
}

func TestFile(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	var f *File
	TestFile(f)

	var b interface{}
	b = f
	if v, ok := b.(ReadWriter); ok {
		fmt.Println(v, ok)
	}
}
