// Date formatting activity

package main

import (
	"time"
	"fmt"
	"os"
)

func parseMonth(month time.Month) int {
	switch month.String() {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	default:
		return 0

	}
}


func main() {
	currentDate := time.Now()
	hour := int(currentDate.Hour())
	minutes := int(currentDate.Minute())
	seconds := int(currentDate.Second())	
	day := int(currentDate.Day())
	month := currentDate.Month()
	monthint := parseMonth(month)
	if monthint == 0 {
		fmt.Println("Invalid month!")
		os.Exit(1)
	}

	year := int(currentDate.Year())

	fmt.Printf("%02d:%02d:%02d %d/%02d/%02d\n", hour, minutes, seconds, year, monthint, day)
}

