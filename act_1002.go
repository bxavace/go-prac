// Enforcing specific date

package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	date := time.Date(2024, 7, 8, 11, 49, 20, 350123, time.UTC)
	day := strconv.Itoa(date.Day())
	month := strconv.Itoa(int(date.Month()))
	year := strconv.Itoa(date.Year())
	hour := strconv.Itoa(date.Hour())
	minute := strconv.Itoa(date.Minute())
	second := strconv.Itoa(date.Second())

	fmt.Printf("%v:%v:%v %v/%v/%v", hour, minute, second, year, month, day)
}
