package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	switch input {
	case "Philip\r\n": // windows换行符\r\n
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goobye!\n")
	}

	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chris\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye\n")
	}

	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! GoodBye!")
	}
}
