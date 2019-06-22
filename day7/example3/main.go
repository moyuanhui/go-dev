package main

//写入文件
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("D:\\output.da", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("an error occurred with file create")
		return
	}

	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
