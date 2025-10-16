package main

import (
	"fmt"
)

func main() {
	responseTimes := []byte{27, 54, 36, 49, 85}

	fmt.Printf("Original array list: %v\n", responseTimes)

	// append to a byte slice 
	update := append(responseTimes, 20)

	// get cap and length of the array
	fmt.Printf("Updated array list: %v, len: %d, cap: %d\n", update, len(update), cap(update))

	// reslice the list
	resliced := update[:3]
	fmt.Printf("Resliced array (first 3 elements): %v, len: %d, cap: %d\n", resliced, len(resliced), cap(resliced))
	
}