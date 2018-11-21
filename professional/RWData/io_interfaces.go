package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// unbuffered
	a, err := fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(a)
	fmt.Printf("the len of 'hello world! - unbuffered' is : %d\n", len("hello world! - unbuffered"))
	// buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)
	// and now so does buf
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush()
}
