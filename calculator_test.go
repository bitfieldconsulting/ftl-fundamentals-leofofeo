package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
	name string
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 1, b: 1, want: 2, name: "Add 1 + 1"},
		{a: 10.0, b: -2, want: 8.0, name: "Add a positive and a negative"},
		{a: 5, b: 3, want: 8, name: "Add 5 + 3"},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Error for the following test: ")
			t.Errorf("%s", tc.name)
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0, name: "Subtract 2 from 2"},
		{a: 0, b: -15, want: 15, name: "Subtract a negative number from 0"},
		{
			a:    10,
			b:    100,
			want: -90,
			name: "Subtract a larger value from a smaller value"},
		{
			a:    10,
			b:    .5,
			want: 9.5,
			name: "Subtract a decimal from a whole number",
		},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Error for the following test: ")
			t.Errorf("%s", tc.name)
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Multiply 2 and 2"},
		{a: 1, b: -1, want: -1, name: "Multiply a positive and a negative"},
		{a: 10, b: 100, want: 1000, name: "Multiply 10 and 100"},
		{
			a:    .5,
			b:    -40,
			want: -20,
			name: "Multiply a positive decimal and a negative integer",
		},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Error for the following test: ")
			t.Errorf("%s", tc.name)
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}
