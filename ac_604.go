// Custom Error Message + Panic + Recovery

package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrLastName = errors.New("invalid last name")
	ErrRoutingNo = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName string
	firstName string
	bankName string
	routingNumber int
	accountNumber int
}

func (d directDeposit) validateRoutingNumber() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if d.routingNumber  < 100 {
		panic(ErrRoutingNo)
	}
	// return nil
}

func (d directDeposit) validateLastName() error {
	if d.lastName == "" || len(d.lastName) <= 0 {
		return ErrLastName
	}
	return nil	
}

func (d directDeposit) report() {
	fmt.Println(strings.Repeat("*", 25))
	fmt.Printf("Last Name:	%v\n", d.lastName)
	fmt.Printf("First Name:	%v\n", d.firstName)
	fmt.Printf("Bank Name: %v\n", d.bankName)
	fmt.Printf("Routing Number:	%v\n", d.routingNumber)
	fmt.Printf("Account Number:	%v\n", d.accountNumber)
}

func main() {
	guy := directDeposit{
		lastName: "",
		firstName: "Abe",
		bankName: "XYZ Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	err := guy.validateLastName()
	if err != nil {
		fmt.Println(err)
	}
	guy.validateRoutingNumber()

	fmt.Println()
	guy.report()
}
