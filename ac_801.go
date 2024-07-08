// Interfaces - Calculating Pay and Performance
package main

import (
	"fmt"
	"errors"
	"os"
)


type Employee struct {
	Id int
	FirstName string
	LastName string
}

type Developer struct {
	Individual Employee
	HourlyRate float64 
	HoursWorkedInYear float64
	Review map[string]interface{}
}

type Manager struct {
	Individual Employee
	Salary float64 
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d Developer) Pay() (string, float64) {
	fullName := d.Individual.FirstName + " " + d.Individual.LastName
	return fullName, (d.HourlyRate * d.HoursWorkedInYear)
}

func (m Manager) Pay() (string, float64) {
	fullName := m.Individual.FirstName + " " + m.Individual.LastName
	return fullName, (m.Salary + (m.Salary * m.CommissionRate))
}

func payDetails(p Payer) {
	name, payment := p.Pay()
	fmt.Println("Name:", name)
	fmt.Println("Payment:", payment)
}

func main() {
	m := make(map[string]interface{})
	m["work quality"] = "excellent"
	m["teamwork"] = 4
	m["skills"] = 3

	e := Employee{500, "Brylle", "Nunez"}
	d := Developer{e, 10, 250, m}
	payDetails(d)
	err := d.reviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseReview(str string) (int, error) {
	switch str {
	case "excellent":
		return 5, nil
	case "good":
		return 4, nil
	case "fair":
		return 3, nil
	case "poor":
		return 2, nil
	case "unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + str)
	}
}

func overallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := parseReview(v)
		if err != nil {
			return 0, err
		}

		return rating, nil
	default:
		return 0, errors.New("unknown type")
	}
}

func (d Developer) reviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := overallReview(v)
		if err != nil {
			return err 
		}
		total += rating
	}
	ave := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.Individual.FirstName, ave)
	return nil
}


