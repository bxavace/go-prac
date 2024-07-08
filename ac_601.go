// Custom Error Message

package main

import (
	"errors"
	"fmt"
)

var (
	ErrLastName = errors.New("invalid last name")
	ErrRoutingNo = errors.New("invalid routing number")
)

func main() {
	fmt.Println(ErrLastName)
	fmt.Println(ErrRoutingNo)
}
