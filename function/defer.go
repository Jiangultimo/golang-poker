package main

import "fmt"

// 关键字defer允许我们推迟到函数返回之前（或任意位置执行return语句之后）一刻才执行某个语句或函数
func function1() {
	fmt.Print("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom\n")
}

func function2() {
	fmt.Printf(("Function2: Deferred until the end of the calling function!"))
}

func main() {
	function1()
}

// print
// In function1 at the top
// In function1 at the botton
// Function2: Deferred until the end of the calling function!


