package main

import (
	"fmt"
)

func main() {
	//Slices
	listRune := [...]rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}
	listRuneManageable := make([]string, 0)
	listInt := make([]int, 0)
	listAsst := make([]string, 0)

	//Variables
	var (
		d string
		e int
	)

	//Convert the orginal slice into a format to compare with output after the two convertions
	for _, i := range listRune {
		d = string(i)
		listRuneManageable = append(listRuneManageable, d)
	}
	// New slice of rune converted to ASCII values
	fmt.Println()
	fmt.Print(" An Initial Slice of ")
	fmt.Println(listRuneManageable)
	fmt.Println()

	// ASCII value for each rune
	for _, i := range listRune {
		d = string(i)
		e = int(i) //This is the convertion
		listInt = append(listInt, e)
		fmt.Printf(("The ASCII value of " + d + " is "))
		fmt.Print(i)
		fmt.Println(".")
	}
	// New slice of rune converted to ASCII values
	fmt.Println()
	fmt.Print(" A ASCII Slice of ")
	fmt.Println(listInt)
	fmt.Println()

	// ASCii value converted rune
	for _, i := range listInt {
		d = string(i) //This is the convertion
		listAsst = append(listAsst, d)
		fmt.Print("The Character of ")
		fmt.Print(d)
		fmt.Print(" has an ASCII value of ")
		fmt.Print(i)
		fmt.Println(".")
	}

	// New slice of ASCII values converted to rune
	fmt.Println()
	fmt.Print(" A Rune Slice of ")
	fmt.Print(listAsst)
	fmt.Println(" like the initial slice")
	fmt.Println()
}
