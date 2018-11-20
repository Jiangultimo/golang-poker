package main

type Foo map[string]string
type Bar struct {
	thingOne string
	THingTwo int
}

func main() {
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).THingTwo = 1

	z := make(Bar) //  编译错误：cannot make type Bar
	(*z).thingOne = "hello"
	(*z).THingTwo = 1

	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	u := new(Foo)
	(*u)["x"] = "goodbyd" // 运行错误！panic: assignment to entry in nil map
	(*u)["y"] = "world"
}
