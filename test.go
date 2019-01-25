package main

import (
	"fmt"
)

func main() {
	// Strings do not need to be declared as a type, <name> := <desired string sequence>;
	// strang := "henlo." shorthand declaration
	var strang string
	strang = "henlo"
	fmt.Printf("%s \n", strang)
	fmt.Printf("%d", pow(2, 22))
}

func pow(x int, power int) int {
	if power < 0 {
		return -1
	}
	prod := 1
	for i := 0; i < power; i++ {
		prod *= x
	}
	return prod
}
