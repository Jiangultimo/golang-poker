package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "0" // 这里相当于在f1作用域中新声明并初始化了变量a
	print(a)
	f2()
}

func f2() {
	print(a) // 打印全局变量
}