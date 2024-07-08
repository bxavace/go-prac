// Removes an item from a slice.
package main

import "fmt"

var els = []string {
	"Good",
	"Good",
	"Bad",
	"Good",
	"Good",
}

func main() {
	for i, v := range els {
		if v == "Bad" {
			els = append(els[:i], els[i+1:]...)
		}
	}

	fmt.Println(els)
}
