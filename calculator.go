package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Add takes a variable number of float64 and returns their sum.
func Add(nums ...float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}

	return sum
}

// Subtract takes a variable number of float64 and returns their difference.
func Subtract(nums ...float64) float64 {
	difference := nums[0]

	for i := 1; i < len(nums); i++ {
		difference -= nums[i]
	}
	return difference
}

// Multiply takes a variable number of float64 and returns their product.
func Multiply(nums ...float64) float64 {
	product := 1.0
	for _, num := range nums {
		product *= num
	}
	return product
}

// Divide takes a variable number of float64 and returns their quotient, or an error if it's not a valid division.
func Divide(nums ...float64) (float64, error) {
	dividend := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0, errors.New("Cannot divide by zero")
		}
		dividend /= nums[i]
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
	// op := strings.ReplaceAll(operation, " ", "")
	s := strings.ReplaceAll(operation, " ", "")
	operators := []string{"+", "-", "*", "/"}

	var result float64
	var err error

	for _, op := range operators {
		if strings.Contains(s, op) {
			s := strings.Split(s, op)
			a, _ := strconv.ParseFloat(s[0], 64)
			b, _ := strconv.ParseFloat(s[1], 64)

			switch op {
			case "+":
				result = Add(a, b)
			case "-":
				result = Subtract(a, b)
			case "*":
				result = Multiply(a, b)
			case "/":
				result, err = Divide(a, b)
				if err != nil {
					return 0, err
				}
			}

		}
	}

	return result, nil
}
