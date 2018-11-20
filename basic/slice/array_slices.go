package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included
	// print init slice1
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("The init slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The capacity of init slice1 is %d\n", cap(slice1))

	// load the array width integers: 0, 1, 2, 3, 4, 5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	// grow the slice beyond capacity
	// slice1 = slice1[0:7] // panic: runtime error: slice bound out of range
	// 绝对不要用指针指向slice。切片本身已经是一个引用类型，所以它本身就是一个指针
}
