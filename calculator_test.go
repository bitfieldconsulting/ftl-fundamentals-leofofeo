package calculator_test

import (
	"calculator"
	"math/rand"
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

func TestAddRandom(t *testing.T) {
	t.Parallel()
	iterations := 100
	for i := 0; i < iterations; i++ {
		a := rand.Float64() * 10
		b := rand.Float64() * 10

		sum := a + b

		got := calculator.Add(a, b)
		if sum != got {
			t.Errorf("Expected sum of %f and %f to match", a, b)
			t.Errorf("Got %f and %f (the latter from Add)", sum, got)
		}
	}
}

func TestSqrRoot(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{
			a:           4,
			want:        2,
			name:        "Square root of 4",
			errExpected: false,
		},
		{
			a:           16,
			want:        4,
			name:        "Square root of 16",
			errExpected: false,
		},
		{
			a:           100,
			want:        10,
			name:        "Square root of 100",
			errExpected: false,
		},
		{
			a:           -20,
			name:        "Square root of -20",
			errExpected: true,
		},
	}
	for _, tc := range testCases {
		got, err := calculator.SqrRoot(tc.a)
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

func TestVariadicFunctions(t *testing.T) {
	t.Parallel()

	// test addition
	got := calculator.Add(1, 2, 3, 4)
	var want float64 = 10
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}

	got = calculator.Add(1, 2, 3, 4, 100, 1000, -5)
	want = 1105
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}

	// test substraction
	got = calculator.Subtract(10, 5, 3)
	want = 2
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}

	got = calculator.Subtract(-10, 5, -20)
	want = 5
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}

	// test multiplication
	got = calculator.Multiply(-5, 6, 2, 4, 0)
	want = 0
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}

	got = calculator.Multiply(10, -10, 1, -1, .5)
	want = 50
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}

	// test division
	got, err := calculator.Divide(100, 10, 10)
	want = 1
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}

	if err != nil {
		t.Errorf("Error when atttempting to divide")
	}

	got, err = calculator.Divide(-5, 10, .5)
	want = -1
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}

	if err != nil {
		t.Errorf("Error when atttempting to divide")
	}

	_, err = calculator.Divide(100, 10, 10, 0, 20)
	if err == nil {
		t.Errorf("No error thrown when atttempting to divide by zero")
	}
}

type stringTestCase struct {
	a    string
	want float64
	name string
}

func TestStringOperations(t *testing.T) {
	t.Parallel()
	testCases := []stringTestCase{
		{a: "1+5", want: 6, name: "1 + 5"},
		{a: "10 + 20", want: 30, name: "10 + 20"},
		{a: "6 - 4", want: 2, name: "6 - 4"},
		{a: "20-2", want: 18, name: "20 - 2"},
		{a: "5 * 3", want: 15, name: "5 * 3"},
		{a: "1*12.5", want: 12.5, name: "1 * 12.5"},
		{a: "10/10", want: 1, name: "10 / 10"},
		{a: "5   /    10", want: .5, name: "5 / 10"},
	}
	for _, tc := range testCases {
		got, err := calculator.StringOperations(tc.a)
		if tc.want != got {
			t.Errorf("Error for the following test: ")
			t.Errorf("%s", tc.name)
			t.Errorf("want %f, got %f", tc.want, got)
		}
		if err != nil {
			t.Errorf("Error for the following test: ")
			t.Errorf("%s", tc.name)
		}
	}
}
