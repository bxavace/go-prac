// Measuring elapsed time

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("execution took %v seconds.\n", elapsed)
}
