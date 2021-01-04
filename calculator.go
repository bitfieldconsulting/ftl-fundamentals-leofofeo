// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(nums ...float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}

	return sum
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(nums ...float64) float64 {
	difference := nums[0]

	for i := 1; i < len(nums); i++ {
		difference -= nums[i]
	}
	return difference
}

// Multiply takes two numbers and returns their product.
func Multiply(nums ...float64) float64 {
	product := 1.0
	for _, num := range nums {
		product *= num
	}
	return product
}

// Divide takes two numbers and returns their quotient, or an error if it's not a valid division.
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

// SqrRoot takes a number and returns its square root provided it has, otherwirse it returns an error.
func SqrRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot get square root of negative number")
	}
	return math.Sqrt(a), nil
}
