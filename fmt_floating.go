package main

import "fmt"

func main() {
	v := 1234.0
	v1 := 1234.6
	v2 := 1234.67
	v3 := 1234.678
	v4 := 1234.6789
	v5 := 1234.67891

	fmt.Printf("%10.0f\n", v) 	// 10 width, 0 precision
	fmt.Printf("%10.1f\n", v1)	// 10 width, 1 precision
	fmt.Printf("%10.2f\n", v2)	// 10 width, 2 precision
	fmt.Printf("%10.3f\n", v3)	// 10 width, 3 precision
	fmt.Printf("%10.4f\n", v4)	// and so on...
	fmt.Printf("%10.5f\n", v5)
}
