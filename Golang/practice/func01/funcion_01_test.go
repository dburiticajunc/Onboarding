package func01_test

import (
	"github.com/stretchr/testify/require"
	"practice/func01"
	"testing"
)

func TestCalculateSalaryTax_higher_150K(t *testing.T) {
	var salary float32 = 150000
	var salaryTax float32 = func01.CalculateSalaryTax(salary)
	var expectedSalatyTax float32 = 40500
	require.Equal(t, expectedSalatyTax, salaryTax, "Validate the tax, salary higher 150K")
}

func TestCalculateSalaryTax_minor_150(t *testing.T) {
	var salary float32 = 100000
	var salaryTax float32 = func01.CalculateSalaryTax(salary)
	var expectedSalatyTax float32 = 17000
	require.Equal(t, expectedSalatyTax, salaryTax, "Validate the tax, salary minor 150K")
}

func TestGradePointAverageCalculator_negative_number(t *testing.T) {
	var expect float32 = 0
	result, err := func01.GradePointAverageCalculator(1, 2, 4, -5)
	require.Equal(t, expect, result)
	require.Error(t, err)
	/*
		if func01.ErrNegativeNumber.Error() = err.Error() {
			t.Errorf("Expected : %s, got : %s", func01.ErrNegativeNumber.Error(), err.Error())
			return
		}
	*/
}

func TestGradePointAverageCalculator_Successful(t *testing.T) {
	var expect float32 = 3
	result, err := func01.GradePointAverageCalculator(2, 3, 4)
	require.Equal(t, expect, result)
	require.NoError(t, err)
}
