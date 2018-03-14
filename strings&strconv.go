package main

import (
	"fmt"
	"strings"
)

func main() {
	/**
	* HasPrefix 判断字符串s是否以prefix开头
	* strings.HasPrefix(s, prefix string) bool
	*
	* HasSuffix判断字符串s是否以suffix结尾
	* strings.HasSuffix(s, suffix string) bool
	* 包含使用Contains
	* strings.Contains(s, substr string) bool
	 */
	var str string = "this is an example of string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "th"))

	/*
		判断子字符串火字符在付字符串中出现的位置（索引）
	*/
	var str2 string = "Hi, I'm Marc, Hi"
	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "Marc"))
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str2, "Hi"))
	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "Burger"))

	//字符替换
	var str3 string = "hello, hing"
	fmt.Printf(strings.Replace(str3, "hing", "world", 5)) // hello, world

	//统计字符串出现次数
	str4 := "Hello, how is it going, Hugo?"
	manyG := "gggggg"
	fmt.Printf("Number os H's in %s is : ", str4)
	fmt.Printf("%d\n", strings.Count(str4, "H"))
	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	//使用strings.TrimSpace(str)来提出字符串开头和结尾的空白符号；剔除指定字符串，可以使用strings.Trim(str, "cut")来将开头和结尾的cut去掉，以及strings.TrimLeft,strings.TrimRight
	//分割字符串，使用strings.Fields(str)将会利用1个或者多个空白符号来作为动态长度的分隔符将字符串庚哥成若干小块，并返回一个slice，以及strings.Split(str, sep)
	//拼接slice到字符串
	str5 := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str5)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("$s -", val)
	}
	fmt.Println()
	str6 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str6, "|")
	fmt.Printf("Splitted in slice:%v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str7 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str7)

	//与字符串相关的类型转换都是通过strconv包来实现
}
