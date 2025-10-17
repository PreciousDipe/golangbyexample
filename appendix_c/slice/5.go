package main

import (
	"fmt"
)

func main() {
	// define the array
	capacity := []byte{2, 4, 5, 9, 3}
	
	data := capacity

	// remove elements from the initial array
	// data = []byte{data[3]}
	data = append(data[:2], data[3:]...)

	// reduce capacity of the array
	new := make([]byte, len(data))
	copy(new, data)

	fmt.Printf("Reduced byte slice: %v, len: %d, cap: %d\n", new, len(new), cap(new))
}