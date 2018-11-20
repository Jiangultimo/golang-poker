package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("fileoutput.dat", os.O_WRONLY|os.O_CREATE, 0666) // 以只写模式打开output.dat，如果文件不存在则自动创建
	if outputError != nil {
		fmt.Printf("An error occurred width file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
