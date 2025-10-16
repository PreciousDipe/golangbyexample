package main

import (
	"fmt"
)

func main() {
	// create an array for the time
	responseTimes := []int{27, 54, 36, 49, 85}

	// loop through the array
	for _, time := range responseTimes {
		// print the response time
		fmt.Printf("Server Response Times %d minutes\n", time)
	}
}