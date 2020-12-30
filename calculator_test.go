package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	name        string
	errExpected bool
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

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{
			a:           2,
			b:           2,
			want:        1,
			name:        "Divide 2 and 2",
			errExpected: false,
		},
		{
			a:           1,
			b:           -1,
			want:        -1,
			name:        "Divide a positive and a negative",
			errExpected: false,
		},
		{
			a:           10,
			b:           100,
			want:        .1,
			name:        "Divide a smaller number by a larger number",
			errExpected: false,
		},
		{
			a:           -40,
			b:           .5,
			want:        -80,
			name:        "Divide a negative integer by a positive decimal",
			errExpected: false,
		},
		{
			a:           6,
			b:           0,
			want:        0,
			name:        "Divide a positive integer by zero",
			errExpected: true,
		},
		{
			a:           0,
			b:           10,
			want:        0,
			name:        "Divide zero by an integer",
			errExpected: false,
		},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if tc.errExpected {
			if err != nil {
				if tc.want != got {
					t.Errorf("Error for the following test: ")
					t.Errorf("%s", tc.name)
					t.Errorf("want %f, got %f", tc.want, got)
				}
			} else {
				t.Errorf("Error for the following test: ")
				t.Errorf("%s", tc.name)
				t.Errorf("Expected error but got none.")
			}
		} else {
			if err != nil {
				t.Errorf("Error for the following test: ")
				t.Errorf("%s", tc.name)
				t.Errorf("Got error - invalid inputs")
			} else {
				if tc.want != got {
					t.Errorf("Error for the following test: ")
					t.Errorf("%s", tc.name)
					t.Errorf("want %f, got %f", tc.want, got)
				}
			}
		}

	}
}
