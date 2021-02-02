package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func TestAddSubtractMultiply(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		a, b, want float64
		extra      []float64
		f          func(float64, float64, ...float64) float64
	}{
		{
			name: "Add 1 + 1",
			a:    1,
			b:    1,
			want: 2,
			f:    calculator.Add,
		},
		{
			name: "Add a positive and a negative",
			a:    10.0,
			b:    -2,
			want: 8.0,
			f:    calculator.Add,
		},
		{
			name:  "Add a whole bunch of numbers",
			a:     5,
			b:     3,
			extra: []float64{2, 1},
			want:  11,
			f:     calculator.Add,
		},
		{
			name: "Subtract 2 from 2",
			a:    2,
			b:    2,
			want: 0,
			f:    calculator.Subtract,
		},
		{
			name:  "Subtract a bunch of negative numbers from 0",
			a:     0,
			b:     -15,
			extra: []float64{-2, -3},
			want:  20,
			f:     calculator.Subtract,
		},
		{
			name: "Subtract a larger value from a smaller value",
			a:    10,
			b:    100,
			want: -90,
			f:    calculator.Subtract,
		},
		{
			name: "Subtract a decimal from a whole number",
			a:    10,
			b:    .5,
			want: 9.5,
			f:    calculator.Subtract,
		},
		{
			a:    2,
			b:    2,
			want: 4,
			name: "Multiply 2 and 2",
			f:    calculator.Multiply,
		},
		{
			a:    1,
			b:    -1,
			want: -1,
			name: "Multiply a positive and a negative",
			f:    calculator.Multiply,
		},
		{
			a:    10,
			b:    100,
			want: 1000,
			name: "Multiply 10 and 100",
			f:    calculator.Multiply,
		},
		{
			name: "Multiply a positive decimal and a negative integer",
			a:    .5,
			b:    -40,
			want: -20,
			f:    calculator.Multiply,
		},
		{
			name:  "Multiply a bunch of numbers",
			a:     2,
			b:     2,
			extra: []float64{3, 4, 5},
			want:  240,
			f:     calculator.Multiply,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.f(tc.a, tc.b, tc.extra...)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		a, b, want  float64
		extra []float64
		errExpected bool
	}{
		{
			name:        "Divide 2 and 2",
			a:           2,
			b:           2,
			want:        1,
			errExpected: false,
		},
		{
			name:        "Divide a positive and a negative",
			a:           1,
			b:           -1,
			want:        -1,
			errExpected: false,
		},
		{
			name:        "Divide a smaller number by a larger number",
			a:           10,
			b:           100,
			want:        .1,
			errExpected: false,
		},
		{
			name:        "Divide a negative integer by a positive decimal",
			a:           -40,
			b:           .5,
			want:        -80,
			errExpected: false,
		},
		{
			name:        "Divide a positive integer by zero",
			a:           6,
			b:           0,
			want:        0,
			errExpected: true,
		},
		{
			name:        "Divide zero by an integer",
			a:           0,
			b:           10,
			want:        0,
			errExpected: false,
		},
		{
			name:        "Divide something by a bunch of numbers including zero",
			a:           4,
			b:           2,
			extra: []float64{1, 0, 1},
			want:        999,
			errExpected: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func (t *testing.T) {
			got, err := calculator.Divide(tc.a, tc.b, tc.extra...)
			errReceived := err != nil
			if tc.errExpected != errReceived {
				t.Fatalf("Unexpected error status: %v", err)
			}
			if !tc.errExpected && tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	iterations := 100
	for i := 0; i < iterations; i++ {
		a := rand.Float64() * 10
		b := rand.Float64() * 10
		want := a + b
		t.Run("Random addition", func(t *testing.T) {
			got := calculator.Add(a, b)
			if want != got {
				t.Errorf("want %f, got %f", want, got)
			}
		})
	}
}

func TestSqrRoot(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		a, want     float64
		errExpected bool
	}{
		{
			name:        "Square root of 4",
			a:           4,
			want:        2,
			errExpected: false,
		},
		{
			a:           16,
			want:        4,
			name:        "Square root of 16",
			errExpected: false,
		},
		{
			name:        "Square root of 100",
			a:           100,
			want:        10,
			errExpected: false,
		},
		{
			name:        "Square root of -20",
			a:           -20,
			errExpected: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Sqrt(tc.a)
			errReceived := err != nil
			if tc.errExpected != errReceived {
				t.Fatalf("Unexpected error status: %v", err)
			}
			if !tc.errExpected && tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestEvaluateExpr(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name, a     string
		want        float64
		errExpected bool
	}{
		{name: "1 + 5", a: "1+5", want: 6, errExpected: false},
		{name: "10 + 20", a: "10 + 20", want: 30, errExpected: false},
		{name: "6 - 4", a: "6 - 4", want: 2, errExpected: false},
		{name: "20 - 2", a: "20-2", want: 18, errExpected: false},
		{name: "5 * 3", a: "5 * 3", want: 15, errExpected: false},
		{name: "1 * 12.5", a: "1*12.5", want: 12.5, errExpected: false},
		{name: "10 / 10", a: "10/10", want: 1, errExpected: false},
		{name: "5 / 10", a: "5   /    10", want: .5, errExpected: false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.EvaluateExpr(tc.a)
			errReceived := err != nil
			if tc.errExpected != errReceived {
				t.Fatalf("Unexpected error status: %v", err)
			}
			if !tc.errExpected && tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}
