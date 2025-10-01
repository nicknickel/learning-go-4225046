package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is a slice since no numbers between []
	// var colors = []string{"Red", "Green", "Blue"}

	// make(what to make, starting length, number of items it can hold)
	// This is more memory efficient
	var colors = make([]string, 0, 3)
	colors = append(colors, "red", "green", "blue")
	fmt.Println(colors)

	// can still add more items to this slice
	colors = append(colors, "purple", "fuschia")
	fmt.Println(colors)

	colors = remove(colors, 1)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
}

// best way to remove an item from slice
func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
