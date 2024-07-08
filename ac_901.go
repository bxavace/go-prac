// Validating SSNs using Logging

package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

var (
	ErrInvalidSSNLength  = errors.New("Invalid SSN Length")
	ErrInvalidSSNNumbers = errors.New("SSN has non-numeric digits")
	ErrInvalidSSNPrefix  = errors.New("SSN prefix has three zeros")
	ErrInvalidDigitPlace = errors.New("SSN starts with 9")
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func checkSSNLength(ssn string) error {
	if len(ssn) != 9 {
		return ErrInvalidSSNLength
	}
	return nil
}

func checkSSNNumbers(ssn string) error {
	_, err := strconv.ParseFloat(ssn, 64)
	if err != nil {
		return ErrInvalidSSNNumbers
	}
	return nil
}

func checkPrefix(ssn string) error {
	if ssn[:3] == "000" {
		return ErrInvalidSSNPrefix
	}
	return nil
}

func checkDigitPlace(ssn string) error {
	if ssn[0] == '9' {
		if ssn[3] == '7' || ssn[3] == '9' {
			return nil
		} else {
			return ErrInvalidDigitPlace
		}
	}
	return nil
}

func processSSN(ssn string) string {
	nodash := strings.ReplaceAll(ssn, "-", "")
	return nodash
}

func main() {
	validateSSN := []string{
		"123-45-6789",
		"012-8-678",
		"000-12-0962",
		"999-33-3333",
		"087-65-4321",
		"123-45-zzzz",
	}
	// newssn := processSSN(validateSSN[len(validateSSN)-1])
	log.Printf("Checking data %#v\n", validateSSN)
	for i := 0; i < len(validateSSN); i++ {
		log.Printf("Validate data %#v %d of %d\n", validateSSN[i], i, len(validateSSN))
		newssn := processSSN((validateSSN[i]))
		err := checkSSNLength(newssn)
		if err != nil {
			log.Printf("the value of %v caused an error: %v", validateSSN[i], err)
		}

		err = checkSSNNumbers(newssn)
		if err != nil {
			log.Printf("the value of %v caused an error: %v", validateSSN[i], err)
		}

		err = checkPrefix(newssn)
		if err != nil {
			log.Printf("the value of %v caused an error: %v", validateSSN[i], err)
		}

		err = checkDigitPlace(newssn)
		if err != nil {
			log.Printf("the value of %v caused an error: %v", validateSSN[i], err)
		}
	}
}
