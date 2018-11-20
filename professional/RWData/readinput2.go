package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader // inputReader是一个指向bufio.Reader的指针
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin) // 将会创建一个读取器，并将其与标准输入绑定，返回的读取器对象提供一个方法ReadString(delim byte)，该方法从输入法中读取内容，知道碰到delim指定的字符，然后将读取到的内容连同delim自负一起放到缓冲区
	fmt.Println("Please enter some input:")
	input, err = inputReader.ReadString('\n') // 遇到 \n 读取结束
	if err == nil {
		fmt.Printf("The input was: %s\n", input	)
	}
}