// in go test files has a suffix of _test.go
// run tests using go test -v
// go test -run TestSimple -v for running a test
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

type testCase struct {
	value    float64
	expected float64
}

func TestSimple(t *testing.T) { // starts with Test and has t *testing.T
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation - %s", err) // will stop execution and print error
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

func TestMany(t *testing.T) {
	testCases := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)

			if err != nil {
				t.Fatal("error")
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("bad value - %f", out)
			}
		})
	}
}

func TestCSV(t *testing.T) {
	file, err := os.Open("sqrt_cases.csv")

	if err != nil {
		t.Fatalf("Error occurred while reading the csv file")
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		t.Fatalf("Error while reading the csv content")
	}

	testCases := []testCase{}
	for _, line := range lines {
		testCase := testCase{}
		if testCase.value, err = strconv.ParseFloat(line[0], 64); err != nil {
			t.Fatalf("Error occurred while reading first coloumn from csv")
		}
		if testCase.expected, err = strconv.ParseFloat(line[1], 64); err != nil {
			t.Fatalf("Error occurred while reading second coloumn from csv")
		}

		testCases = append(testCases, testCase)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)

			if err != nil {
				t.Fatal("error")
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("bad value - %f", out)
			}
		})
	}
}

// go test -v -bench . will run all the benchmark along with test
// go test -v -bench . -run TTT will run only benchmark
// go test -v -bench . -run TTT -cpuprofile=prof.out
// go tool pprof prof.out
// list <function name>
func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatal(err)
		}
	}
}
