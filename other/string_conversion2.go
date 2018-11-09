package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC"
	var newS string

	fmt.Printf("The Size of ints is : %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting width error\n", orig)
		return
	}
	fmt.Printf("The iterger is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is : %s\n", newS)
}
