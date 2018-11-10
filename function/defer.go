package main

import "fmt"

func function1() {
	fmt.Print("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom\n")
}

func function2() {
	fmt.Printf(("Function2: Deferred until the end of the calling function!"))
}

func main() {
	function1()
}


