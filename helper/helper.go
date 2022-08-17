package helper

import "time"

func Today() time.Time {
	date := time.Now().Format("02/01/2006")
	parsedTodayDate, _ := time.Parse("02/01/2006", date)

	return parsedTodayDate
}

func Parser(date string) time.Time {
	stringDate, _ := time.Parse("02/01/2006", date)

	return stringDate
}
