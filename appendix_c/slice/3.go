package main

import (
	"fmt"
)

func main() {
	responseTimes := []int{27, 54, 36}
	responseTimes1 := []int{49, 85}

	// append the first array to the second array
	result := append(responseTimes, responseTimes1...)
	// return the result
	fmt.Println(result)
}