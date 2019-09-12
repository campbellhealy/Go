// # Penny Calculator
// #
// # Year start is a fixed date to 01/01/2019
// #
// # Calculates how much you will save by the end of the year if you start on that day
// # Start on that day with the number of pennies for that day is calculated
// # Increasing by a penny each day
// #
package main

import (
	"fmt"
	"time"
)

func Date(day, month, year int) time.Time {
	return time.Date(day, time.Month(month), year, 0, 0, 0, 0, time.UTC)
}

func main() {
	// To Accommodate Leap Year - Not used as yet
	YearOne := Date(2019, 1, 1)
	//YearTwo := Date(2019, 12, 31)

	// Calculate Day values required
	dayNow := time.Now()
	days := dayNow.Sub(YearOne).Hours() / 24
	DayNum := int(days)
	YearDays := 365
	leftDays := (YearDays - DayNum)
	TotDays := DayNum

	// Base Output
	fmt.Println("Number of Days in the Year: ", YearDays)
	fmt.Println("Days left in Year: ", leftDays)
	fmt.Println("Day # of Year: ", DayNum)
	fmt.Println()
	// Calculation
	for DayNum < YearDays {
		DayNum++
		TotDays = TotDays + DayNum
	}
	iTotDays := float32(TotDays)
	iTotDays = iTotDays / 100
	fmt.Println("Amount if you start saving today: ", "£", iTotDays)
}
