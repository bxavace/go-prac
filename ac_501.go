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

func (d* Developer) HoursWorked() string {
	m := ""
	total := 0
	for i := 0; i < 7; i++ {
		total += d.WorkWeek[i]
		if d.WorkWeek[i] != 0 {
			m += fmt.Sprintf("Hours worked on %v: %v\n", wToString(i), d.WorkWeek[i])
		}
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
	dev.HourlyRate = 12
	
	dev.LogHours(2, 8)	
	dev.LogHours(1, 10)
	dev.LogHours(3, 4)
	// Should be 8 + 10 + 4 = 22
	fmt.Print(dev.HoursWorked())
}
