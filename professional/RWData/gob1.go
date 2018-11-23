package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	// 初始化encoder和decoder
	// Normally enc and dec would be bound to network connections and the encoder and decoder would run in defferent processes
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // will write to network
	dec := gob.NewDecoder(&network) // will read from network
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}
