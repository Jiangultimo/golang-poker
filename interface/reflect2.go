package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v: ", v.CanSet())
	v = reflect.ValueOf(&x) // take the address of x
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.14159)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
