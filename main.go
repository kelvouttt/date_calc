package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now().Format("2006-01-02")
	fmt.Printf("Today is %v\n", currentTime)

	// Finding the dates difference between 2 point of dates
	differenceDur := differenceDate()
	fmt.Printf("The difference between the 2 dates are %v days or %v years.\n",
		int(differenceDur.Hours())/24,
		int(differenceDur.Hours())/24/12
		int(differenceDur.Hours()/24/365))

	// Adding/subtracting days, months, years
	updatedDate := newDate()
	fmt.Printf("It's %v", updatedDate)
}

// This function does not take any arguments, though it will prompt a user input of 2 different dates.
// The first date will typically be the t1 and second as t2. To find the duration, we will subtract the 2 dates t2 - t1 using the time.Sub() function
func differenceDate() time.Duration {
	var t1 string
	var t2 string

	fmt.Println("Enter the first date to find the difference:")
	fmt.Scan(&t1)
	fmt.Println("Enter the second date:")
	fmt.Scan(&t2)

	firstDate, _ := time.Parse("02/01/2006", t1)
	secondDate, _ := time.Parse("02/01/2006", t2)
	difference := secondDate.Sub(firstDate)

	return difference
}

func newDate() string {
	var dates string
	var year int
	var month int
	var day int

	fmt.Println("Enter days for calculation")
	fmt.Scan(&day)
	fmt.Println("Enter month for calculation")
	fmt.Scan(&month)
	fmt.Println("Enter year for calculation")
	fmt.Scan(&year)
	fmt.Println("Please enter a date:")
	fmt.Scan(&dates)
		

	parsedDate, _ := time.Parse("02/01/2006", dates)
	parsedNewDate := parsedDate.AddDate(year, month, day).Format("02/01/2006")

	return parsedNewDate
}
