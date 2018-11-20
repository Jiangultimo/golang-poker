package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is :%d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is : %c \n", pos ,char)
	}
	fmt.Println()
	str2 := "Chinese: 日本语"
	fmt.Printf("The length of str2 is : %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("charater %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune char byte" )
	for index, rune := range str2 {
		fmt.Printf("%- 2d %d %U '%c % X\n", index, rune, rune, rune,[]byte(string(rune)))
	}
	practivce()
}

func practivce() {
	for i := 0; i < 5; i++{
		var v int // 初始化零值 0
		fmt.Printf("%d", v)
		v = 5
	}
}


