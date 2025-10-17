package main

import (
	"fmt"
)

func main() {
	// assign different variables
	a := 45
	b := 29
	c := 50

	// Declare pointers to variables
	varA(&a)
	fmt.Println(a)

	varB(&b)
	fmt.Println(b)

	varC(&c)
	fmt.Println(c)

}

// Mutate the variables
func varA(a *int) {
	*a++
}

func varB(b *int) {
	*b++
}

func varC(c *int) {
	*c++
}