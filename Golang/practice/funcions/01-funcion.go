package main

import (
	"fmt"
)

const (
	minimun = "minimun"
	avarage = "avarage"
	maximun = "maximun"
)

func main() {
	//CalculateSalaryTax(50000)

	//GradePointAverageCalculator(2, 3, 5)

	//SalaryCalculator(14400, "C")

	//Get minor value of a slice
	operMin := CalculateStatistics(minimun)
	resultMin := operMin(1, 2, 3, 4, 5)
	println("The minor value is : ", resultMin)

	//Get bigger value of a slice
	operMax := CalculateStatistics(maximun)
	resultMax := operMax(1, 2, 3, 4, 5)
	println("The bigger value is : ", resultMax)

	//Get Avg of a slice
	operAvg := CalculateStatistics(avarage)
	resultAvg := operAvg(1, 2, 3, 4, 5)
	println("Then avarage of slice is: ", resultAvg)

}

func CalculateSalaryTax(salary float32) (result float32) {
	if salary >= 150000 {
		result = salary * 0.27
	}
	result = salary * 0.17
	fmt.Println("the tax for a salary of ", salary, " is: ", result)
	return result
}

func GradePointAverageCalculator(grades ...float32) (result float32) {
	var counter float32
	var avg float32
	for _, grade := range grades {
		if grade < 0 {
			panic("You cannot enter negative notes")
		}
		counter += grade
	}
	avg = counter / float32(len(grades))
	println("The Avg is: ", avg)
	return counter
}

func SalaryCalculator(min int16, category string) (result float32) {
	var hours float32 = float32(min) / 60.0
	var salaryTem float32
	var hoursOfWorkPerDay float32 = 8.0
	var dayOfWorkPerMounth float32 = 30.0
	switch category {
	case "A":
		var salaryPerHour float32 = 3000.0
		haltSalary := hoursOfWorkPerDay * dayOfWorkPerMounth * salaryPerHour * 0.5
		salaryTem = hours*salaryPerHour + haltSalary
	case "B":
		var salaryPerHour float32 = 1500.0
		haltSalary := hoursOfWorkPerDay * dayOfWorkPerMounth * salaryPerHour * 0.2
		salaryTem = hours*salaryPerHour + haltSalary
	case "C":
		var salaryPerHour float32 = 1000.0
		salaryTem = hours * salaryPerHour
	default:
		panic("The Category is invalid")
	}
	println("The salary for category ", category, "with ", hours, "hours is: $", int32(salaryTem))
	return salaryTem
}

func CalculateStatistics(operation string) func(n ...float64) float64 {

	switch operation {
	case minimun:
		return MinValue
	case maximun:
		return MaxValue
	case avarage:
		return AvgValue
	}
	return nil
}

func MinValue(n ...float64) (result float64) {
	minorNumber := n[0]
	for _, number := range n {
		if number < minorNumber {
			minorNumber = number
		}
	}
	result = minorNumber
	return
}
func MaxValue(n ...float64) (result float64) {
	biggerNumber := n[0]
	for _, number := range n {
		if number > biggerNumber {
			biggerNumber = number
		}
	}
	result = biggerNumber
	return
}
func AvgValue(n ...float64) (result float64) {
	counter := 0.0
	for _, number := range n {
		counter += number
	}
	return counter / float64(len(n))
}
