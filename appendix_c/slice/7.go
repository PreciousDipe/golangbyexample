package main

import (
	"fmt"
)

func main() {
	// define the array
	array := []int{2, 4, 6, 7, 3, 2, 4}
	// set the index to 0
	index := 0

	// loop through the array
	for i := 0; i < len(array); i++ {
		// checks if duplicates exists
		isDuplicate := false
		
		// checks if the next element is the same the the previous element
		for j :=0; j < index; j++ {
			if array[i] == array[j] {
				isDuplicate = true
				break
			}
		}
		
		// if not moves to the next element
		if !isDuplicate {
		array[index] = array[i]
		index++
	}
	}

	fmt.Println(array[:index])
}