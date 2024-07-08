// Timezones

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	LATime, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("Error in getting LA Time.")
	}

	NYTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error in getting NY Time.")
	}
	fmt.Printf("The local current time is: %v\n", currentTime.Format(time.ANSIC))
	fmt.Printf("The time in New York is: %v\n", currentTime.In(NYTime).Format(time.ANSIC))
	fmt.Printf("The time in Los Angeles is: %v\n", currentTime.In(LATime).Format(time.ANSIC))
}
