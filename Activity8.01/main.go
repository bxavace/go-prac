package main

import (
	"fmt"
	pr "Activity8.01/payroll"
	"os"
)

var employeeReview = make(map[string]interface{})

func init() {
	fmt.Println("Employee Pay and Performance Review")
	fmt.Println("+++++++++++++++++++++++++++++++++++")
}

func init() {
	fmt.Println("Initializing variables")
	employeeReview["WorkQuality"] = 5
	employeeReview["Teamwork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
}

func main() {
	d := pr.Developer{Individual: pr.Employee {Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := pr.Manager{Individual: pr.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pr.PayDetails(d)
	pr.PayDetails(m)
}
