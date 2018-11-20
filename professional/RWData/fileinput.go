package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("fileinput.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close() // 在程序最后的时候关掉os.Open()?

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n') // 在使用ReadString和ReadBytes方法的时候，不需要关心操作系统的类型，直接使用\n就行了
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF { // 常量io.EOF的值是true
			return
		}
	}
}
