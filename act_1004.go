// Future date and time

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("The current time: %v\n", now.Format(time.ANSIC))
	addedTime := now.Add(time.Second * 6)
	addedTime = addedTime.Add(time.Minute * 6)
	addedTime = addedTime.Add(time.Hour * 6)
	fmt.Printf("6 hours, 6 minutes, 6 seconds from now: %v", addedTime.Format(time.ANSIC))
}
