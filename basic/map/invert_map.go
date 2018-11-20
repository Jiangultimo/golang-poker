package main

import "fmt"

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
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal { // 对调key，value
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v /", k, v)
	}
}
