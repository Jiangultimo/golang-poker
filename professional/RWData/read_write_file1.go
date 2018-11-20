package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "read_write_file1_input.txt"
	outputFile := "read_write_file1_output.txt"
	buf, err := ioutil.ReadFile(inputFile) // 第一个返回值的类型是[]byte，里面存放读取到的内容，第二个返回值是错误
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // 函数WriteFile()可以将[]byte的值写入文件， oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
