package main

import "fmt"

func main() {
	seasons := []string{"Spring", "Sumber", "Autum", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is $s\n", ix, season)
	}
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
}
