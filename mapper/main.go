package main

import "fmt"

type Results map[int][]Employee

type Employee struct {
	Name      string
	Role      string
	Pay       float64
	Seniority int
}

func main() {
	employees := []Employee{
		{"Alice Johnson", "Software Engineer", 75000, 2},
		{"Bob Smith", "DevOps Engineer", 85000, 3},
		{"Carol Brown", "Project Manager", 95000, 3},
		{"David Lee", "Backend Developer", 75000, 1},
		{"Emma Wilson", "Frontend Developer", 75000, 2},
		{"Frank Harris", "Team Lead", 95000, 3},
		{"Grace Adams", "Product Manager", 85000, 3},
		{"Henry Clark", "Quality Analyst", 68000, 1},
		{"Ivy Hall", "UI/UX Designer", 68000, 2},
		{"Jack Miller", "Tech Support Engineer", 68000, 1},
	}
	mapper := make(Results)
	// group data with levels
	for _, employee := range employees {
		mapper[employee.Seniority] = append(mapper[employee.Seniority], employee)
	}
	fmt.Println(mapper)

}
