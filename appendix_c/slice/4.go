package main

import (
	"fmt"
)

func main() {
	// define the array
	database := []int{3, 7, 5, 10, 15}
	// get the total sum of the array
	totaldata := dataSum(database)
	fmt.Println(totaldata)
}

// create a sum function to add the array together
func dataSum(data []int) int {
	// start the count from 0
	counts := 0
	// loop through the array
	for i, sum := range data {
		// add them up
		counts += sum
		fmt.Println(i, sum)
	}
	return counts
}

