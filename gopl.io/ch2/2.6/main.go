package main

import (
	"fmt"
	"os"
	"package-base/popcount"
	"package-base/tempconv"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Print(popcount.PopCount(45))
	}
}
