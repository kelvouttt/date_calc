package main

import (
	"fmt"
	"time"

	"github.com/spf13/calc/helper"
)

func main() {
	today := helper.Today()
	fmt.Printf("Today is %v\n", today.Format("02/01/2006"))

	for {
		input := userInput()

		if input == 1 {
			// Finding the dates difference between 2 point of dates
			differenceDur := differenceDate()
			fmt.Printf("The difference between the 2 dates are %v days or %v months or %v year(s).\n",
				int(differenceDur.Hours())/24,
				int(differenceDur.Hours())/24/30,
				int(differenceDur.Hours()/24/365))
		} else if input == 2 {
			// Adding/subtracting days, months, years
			updatedDate := newDate()
			fmt.Printf("It's %v\n", updatedDate)
		} else if input == 3 {
			break
		} else {
			fmt.Println("Please choose option 1 or 2 to start.")
		}
	}
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

	firstDate := helper.Parser(t1)
	secondDate := helper.Parser(t2)
	difference := secondDate.Sub(firstDate)

	return difference
}

func newDate() string {
	var dates string
	var year int
	var month int
	var day int

	fmt.Println("Please enter a date:")
	fmt.Scan(&dates)
	fmt.Println("Enter days for calculation")
	fmt.Scan(&day)
	fmt.Println("Enter month for calculation")
	fmt.Scan(&month)
	fmt.Println("Enter year for calculation")
	fmt.Scan(&year)

	parsedDate := helper.Parser(dates)
	parsedNewDate := parsedDate.AddDate(year, month, day).Format("02/01/2006")

	return parsedNewDate
}

func userInput() int {
	var input int

	fmt.Println("Enter the number to start using the program or to exit the program.")
	fmt.Println("1. Calculating date difference")
	fmt.Println("2. Adding/subtracting date")
	fmt.Println("3. Exit")
	fmt.Println("Input:")
	fmt.Scan(&input)

	return input
}
