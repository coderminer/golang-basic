package main

import (
	"fmt"
)

type Salary struct {
	Basic, HRA, TA float64
}
type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("=================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "-------------------"
}

func main() {

	e := Employee{
		FirstName: "Kevin",
		LastName:  "Woo",
		Email:     "test@mail.com",
		Age:       12,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.0,
				TA:    2000.0,
			},
			Salary{
				Basic: 16000.0,
				HRA:   6000.0,
				TA:    2100.0,
			},
		},
	}

	fmt.Println(e.EmpInfo())
}
