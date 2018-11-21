package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') // 循环读取，以 \n 换行符为标识
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 { // 如果没有参数，则直接打印控制台输入的内容
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ { // 遍历参数
		f, err := os.Open(flag.Arg(i)) // 打开参数中的文件
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s:%s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}