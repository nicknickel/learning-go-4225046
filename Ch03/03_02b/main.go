package main

import (
	"fmt"
)

func main() {

	anInt := 42
	// * is the "dereferencing" operator
	// & is memory address of variable
	var p *int = &anInt

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("value of p:", *p)
	}

	value1 := 42.13
	pointer1 := &value1
	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer:", *pointer1)
	fmt.Println("Original Value:", value1)
}
