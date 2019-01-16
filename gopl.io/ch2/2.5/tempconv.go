package main

import "fmt"

type Celsius float64 // 摄氏度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC Celsius = 0 // 结冰点温度
	BoilingC Celsius = 100 // 沸水温度
)

// 传入摄氏度，返回华氏温度
// 类型转换，类型转换不会改变值本身，但是会使他们的语义发生变化
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

// 传入华氏温度，返回设置温度
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC - FreezingC)
	boilingF := CtoF(BoilingC)
	fmt.Printf("%g\n", boilingF - CtoF(FreezingC))
	// 不同类型不能直接进行操错
	fmt.Printf("%g\n", boilingF - FreezingC) // compile error: type mismatch
}


