package main
/*
	This is my version of the FizzBuzz challenge in go.
	I added in user interaction to give the first and last number of a range.
	As most challenge solutions just add in fixed values, that are not useful as a tool
	The if statement is to catch and use whichever number is higher
*/

import (
	"fmt"
)

func main(){
	// declaring the variable using the var keyword
	var firstNumber, lastNumber int
	fmt.Println("Please enter the first number of the range")
	fmt.Scanln(&firstNumber)
	fmt.Println("Please enter the last number of the range")
	fmt.Scanln(&lastNumber)

	if firstNumber <= lastNumber {
		for i := firstNumber; i<= lastNumber; i++{
			fizzbuzz(i)}
	} else 	if firstNumber > lastNumber {
		for i := lastNumber; i<= firstNumber; i++{
				fizzbuzz(i)}
	}	
}

func fizzbuzz(i int){
	fizz := "fizz"
	buzz := "buzz"
	fizzbuzz := "fizzbuzz"
	if i%3 == 0 && i%5 ==0 {
		fmt.Println(i, fizzbuzz)
	} else if i%3 == 0 {
		fmt.Println(i, fizz)
	} else if i%5 == 0 {
		fmt.Println(i, buzz)
	} else {
		fmt.Println(i)
	}
}