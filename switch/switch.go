package main

import "fmt"

func main() {
	var num1 int = 7
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default: fmt.Println("Number is 1- or greater")
	}
	 switch0()
}

func switch0() {
	k := 6
	switch k {
	case 4: fmt.Println("was <= 4")
	fallthrough
	case 5:
		fmt.Println("was <= 5")
	fallthrough
	case 6:
		fmt.Println("was <= 6")
	fallthrough // 加了fallthrough, 会执行到default
	default:
		fmt.Println("default case")
	}
	// was <= 6
	// default case
}
