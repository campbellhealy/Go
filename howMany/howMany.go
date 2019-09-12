package main

// The how_many function determine if it's one of or more than 1
// Giving the correct English sentence

import (
	"fmt"
	"strconv"
)

type Foo struct {
	Number int
	Text   string
}

func main() {
	array := []Foo{{Number: 1, Text: "pie"}, {Number: 10, Text: "fish"}, {Number: 5, Text: "apples"}}
	for i := 0; i <= 2; i++ {
		n := strconv.Itoa(array[i].Number)
		g := (array[i].Text)
		if array[i].Number == 1 {
			fmt.Println("There is " + n + " " + g + ".")
		} else {
			fmt.Println("There are " + n + " " + g + ".")
		}
	}
}
