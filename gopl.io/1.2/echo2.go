package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // 每次循环迭代，range都会产生一对值：索引以及在该索引处的元素值
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
