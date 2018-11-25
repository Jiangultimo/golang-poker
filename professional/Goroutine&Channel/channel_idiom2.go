package main

import (
	"fmt"
	"time"
)

func main() {
	suck(pump()) // 通道迭代模式；生产者消费者模式
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
