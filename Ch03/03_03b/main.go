package main

import (
	"fmt"
)

func main() {
	var colors [3]string

	// Notice this isn't using the :
	// because the type is already indicated
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	fmt.Println("Number of colores:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}
