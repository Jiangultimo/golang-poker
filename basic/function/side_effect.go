package main

import "fmt"

func Multiply(a, b int, reply *int) {
	fmt.Println(*reply) // 指向 reply 这个地址，即为 0
	*reply = a * b // 相当于重新为 n 赋值
}

func main() {
	n := 0
	reply := &n // 取n的地址
	Multiply(10, 6, reply) // 这里传进去的是n的地址，直接改变reply的值，这里不再需要return，按引用传递
	fmt.Println("Multiply: ", *reply)
}
