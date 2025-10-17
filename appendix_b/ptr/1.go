package main

import (
	"fmt"
)

func main() {
	// assign different variables
	a := 45
	b := 29
	c := 50

	// gets the pointer of a and prints whats there
	fmt.Println(a)
	varA(&a)
	fmt.Println(a)

	// gets the element in b and prints whats there
	varB(b)
	fmt.Println(b)

	// gets the element in c and prints out the pointer
	varC(c)
	fmt.Println(&c)

}

// updates the element in the pointer of a
func varA(a *int) {
	*a++
}

// updates the element in b locally
func varB(b int) {
	b++
}

//  updates the element in c locally
func varC(c int) {
	c++
}