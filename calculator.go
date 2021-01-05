package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Add takes a variable number of float64 and returns their sum.
func Add(a float64, b float64, nums ...float64) float64 {
	sum := a + b

	if len(nums) > 0 {
		for _, n := range nums {
			sum += n
		}
	}

	return sum
}

// Subtract takes a variable number of float64 and returns their difference.
func Subtract(a float64, b float64, nums ...float64) float64 {
	difference := a - b

	if len(nums) > 0 {
		for _, n := range nums {
			difference -= n
		}
	}

	return difference
}

// Multiply takes a variable number of float64 and returns their product.
func Multiply(a float64, b float64, nums ...float64) float64 {
	product := a * b

	if len(nums) > 0 {
		for _, n := range nums {
			product *= n
		}
	}

	return product
}

// Divide takes a variable number of float64 and returns their quotient, or an error if it's not a valid division.
func Divide(a float64, b float64, nums ...float64) (float64, error) {
	if b == 0 || containsFloat64(nums, 0.0) {
		return 0, errors.New("Cannot divide by zero")
	}

	dividend := a / b

	if len(nums) > 0 {
		for _, n := range nums {
			dividend /= n
		}
	}

	return dividend, nil
}

// SqrRoot takes a number and returns its square root provided it has one, otherwise it returns an error.
func SqrRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot get square root of negative number")
	}
	return math.Sqrt(a), nil
}

// StringOperations takes in an expression as a string and parses it to evaluate it. Ex: "10 + 5" returns 15
func StringOperations(operation string) (float64, error) {
	s := strings.ReplaceAll(operation, " ", "")
	operators := []string{"+", "-", "*", "/"}

	for _, op := range operators {
		if strings.Contains(s, op) {
			s := strings.Split(s, op)
			a, _ := strconv.ParseFloat(s[0], 64)
			b, _ := strconv.ParseFloat(s[1], 64)

			switch op {
			case "+":
				return Add(a, b), nil
			case "-":
				return Subtract(a, b), nil
			case "*":
				return Multiply(a, b), nil
			case "/":
				return Divide(a, b)
			}

		}
	}
	return 0.0, errors.New("Invalid operator")
}

func containsFloat64(s []float64, i float64) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}
