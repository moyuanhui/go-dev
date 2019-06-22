package main

//拷贝文件
import (
	"fmt"
	"io"
	"os"
)

func main() {

	copyFile("D:\\go.gz1", "D:\\go.gz")
	fmt.Println("Copy done!")
}

func copyFile(distName, srcName string) (written int64, error error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	defer src.Close()
	dst, err := os.OpenFile(distName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)

}
