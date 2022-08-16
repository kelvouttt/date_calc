package main

import (
	"fmt"
	"time"
)

func main() {

	date := time.Date(2022, 8, 2, 0, 0, 0, 0, time.Local)

	currentTime := time.Now().Format("2006-01-02")
	fmt.Printf("Today is %v and the type is %T\n", currentTime, currentTime)

	// From the print statement below, can be observed that time.Format returns string data type
	t, _ := time.Parse("2006-01-02", currentTime)
	fmt.Printf("Parsed time value: %v and type %T\n", t, t)

	// Adding/subtracting days, months, years
	fmt.Printf("One day later %v with type %T\n",
		t.AddDate(1, 1, 1).Format("2006-01-02"), t)

	// Finding the dates difference between 2 point of dates
	difference := t.Sub(date)
	fmt.Printf("Days difference between t1: %v and t2: %v with difference of %v days and type %T.\n",
		date,
		t,
		int64(difference.Hours()/24),
		date)

	differenceDur := differenceDate()
	fmt.Printf("The difference between the 2 dates are %v days or %v years.\n",
		int(differenceDur.Hours())/24,
		int(differenceDur.Hours()/24/365))
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
