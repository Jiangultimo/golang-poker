package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	for i := 1; i < 5; i++ {
		map1[i] = float32(i)
	}
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
}