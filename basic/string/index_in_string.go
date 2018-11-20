package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var str string = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is : ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is : ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is : ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is : ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	fmt.Println()

	countString()
	reaptString()
	joinString()
	conversionString()
}

func countString() { // 统计字符串出现次数
	str := "Hello, how is it going, Hugo?"
	manyG := "gggggggggg"
	fmt.Printf("Number of H's in %s is : ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}

func reaptString() { // 重复字符串
	origS := "Hi there!"
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}

func joinString() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}

func conversionString() {
	orig := "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is : %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The intrger is: %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is : %s\n", newS)
}