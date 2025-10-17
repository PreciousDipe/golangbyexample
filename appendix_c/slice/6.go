package main

import (
	"fmt"
)

func main() {
	// define the array
	data := []byte("updated golang version to v:1.25.2")

	// modify the array
	new := make([]byte, len(data))
	copy(new, data)
	new[len(data)-1]++

	// revert the array to a string
	fmt.Println(string(new))
}