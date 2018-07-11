package main 

import (
	"fmt"
	"math"
	"errors"
)

var num int = 10
var numx2, numx3 int

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	multipleReturn()
	errorReturnVal()
	varnumpar()
}

func MultiPly3Nums(a int, b int, c int) int {
	return a * b * c
}

// 命名的返回值
func multipleReturn() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int)(int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int)(x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

// practice
func multReturnVal(a int, b int)(sum int, multiplied int, difference int) {
	sum = a + b
	multiplied = a * b
	difference = a - b
	return
}

func errorReturnVal() {
	fmt.Printf("First example with -1: ")
	ret1, err1:= MySqrt(-1)
	if err1 != nil {
		fmt.Printf("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Printf("It's OK! Return values are: ", ret1, err1)
	}

	fmt.Printf("Second example width 5:")
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Printf("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Printf("It's ok! Return values are: ", ret2, err2)
	}
	fmt.Println(MySqrt2(5))
}

func MySqrt(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(f), nil
}

func MySqrt2(f float64)(ret float64, err error) {
	if f < 0 {
		ret = float64(math.NaN())
		err = errors.New("I wont be able to do a sqrt of negative number!")
	} else {
		ret = math.Sqrt(f)
	}
	return
}

// 传递变长参数
func varnumpar() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimun is: %d\n", x)
	slice := [] int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimun in the slice is : %d", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}