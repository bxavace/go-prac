package main

import "fmt"

type Employee struct {
	Id int
	FirstName string
	LastName string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func wToString(i int) string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[i]
}

func (d* Developer) LogHours(w Weekday, h int) {
	d.WorkWeek[w] = h
}

// Continuation
func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func (d* Developer) PayDay() (int, bool) {
	hours := 0
	for i := 0; i < 7; i++ {
		hours += d.WorkWeek[i]
	}
	payment := 0
	if hours > 40 {
		hours -= 40
		payment = (hours * 2 * d.HourlyRate) + (40 * d.HourlyRate)
		return payment, true 
	} else {
		return hours * d.HourlyRate, false
	}
}

func (d* Developer) PayDetails() {
	for i := 0; i < 7; i++ {
		fmt.Printf("%v hours: %v\n", wToString(i), d.WorkWeek[i])
	}

	weekPay, isOvertime := d.PayDay()
	fmt.Printf("Pay for the week: $ %v\n", weekPay)
	fmt.Printf("Is this overtime pay: %v\n", isOvertime)
}


func (d* Developer) HoursWorked() string {
	m := ""
	total := 0
	for i := 0; i < 7; i++ {
		total += d.WorkWeek[i]
	}
	m += fmt.Sprintf("Hours worked this week: %v\n", total)
	return m
}

func main() {
	var emp Employee
	emp.Id = 123
	emp.FirstName = "Brylle"
	emp.LastName = "Nunez"

	var dev Developer
	dev.Individual = emp
	dev.HourlyRate = 10
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today:", x(2))
	fmt.Println("Tracking hours worked thus far today:", x(3))
	fmt.Println("Tracking hours worked thus far today:", x(5))
	fmt.Println()
	dev.LogHours(Monday, 8)	
	dev.LogHours(Tuesday, 10)
	dev.LogHours(Wednesday, 10)
	dev.LogHours(Thursday, 10)
	dev.LogHours(Friday, 6)
	dev.LogHours(Saturday, 8)
	fmt.Print(dev.HoursWorked())
	dev.PayDetails()
}
