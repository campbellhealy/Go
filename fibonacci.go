package main
// Fib generates fibonacci sequence for n numbers

import "fmt"

func Fib(seqCount int) chan int {
	// Make a channel
	c := make(chan int)
	// raise a function that generates the fib sequence
	go func() {
		a, b := 0, 1
		for i := 0; i < seqCount; i++ {
			a, b = b, a+b
			// send next value to channel
			c <- a
		}
		close(c)
	}()
	return c
}

func main() {
	// Get user interaction
	var seqCount int
	fmt.Println("Please enter the number of sequence values")
	fmt.Scanln(&seqCount)

	// Iterate over the values received by channel
	// every time a new value has arrived, a loop is executed
	for x := range Fib(seqCount) {
		fmt.Println(x)
	}
}
