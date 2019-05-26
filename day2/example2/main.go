package main

import (
	"fmt"
	"os"
)

func main() {
	goos := os.Getenv("GOOS")
	fmt.Println("The operating is:",goos)

	path := os.Getenv("PATH")
	fmt.Println("Path: ",path)
}