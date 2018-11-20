package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int {
		"alpha": 34,
		"bravo": 56,
		"charlie": 23,
		"delta": 87,
		"hotel": 16,
		"indio": 87,
		"juliet": 65,
		"kili": 43,
		"lima": 98,
	}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal)) // 声明一个切片
	i := 0
	for k, _ := range barVal { // 将map里面的值存入切片
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
