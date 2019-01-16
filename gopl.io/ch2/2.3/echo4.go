package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(*n)
	fmt.Println(*sep)
	fmt.Println(flag.Args())
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
