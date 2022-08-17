package helper

import "time"

func TodayDate() time.Time {
	date := time.Now().Format("02/01/2006")
	parsedTodayDate, _ := time.Parse("02/01/2006", date)

	return parsedTodayDate
}
