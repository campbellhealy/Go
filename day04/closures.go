package main

import "fmt"

func intSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Get user interaction
	var seqCount int
	fmt.Println("Please enter the number of sequence values")
	fmt.Scanln(&seqCount)

	nextInt := intSequence()
	for x := 0; x <seqCount; x++ {
		fmt.Println(nextInt())
	} 
	fmt.Println("Finished")
	}
