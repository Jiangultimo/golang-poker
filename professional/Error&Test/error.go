package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var errNotFound error = errors.New("Not Found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
}
