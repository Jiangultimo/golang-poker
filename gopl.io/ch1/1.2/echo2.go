package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // 每次循环迭代，range都会产生一对值：索引以及在该索引处的元素值
		s += sep + arg // ++和--都只能放在变量后面
		sep = " "
	}
	/**
	for循环的形式
	for initialization; condition; post {}
	initialization可选，condition是一个布尔表达式，post语句在循环执行结束后执行，之后再次对condition求值。condition值为false时，循环结束。
	 */
	fmt.Println(s)
}
