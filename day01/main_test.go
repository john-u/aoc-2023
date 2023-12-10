package main

import (
	"testing"
)

const ExampleInput string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestFindCalibrationValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "1abc2",
			want:  12,
		},
		{
			input: "a1b2c3d4e5f",
			want:  15,
		},
		{
			input: "treb7uchet",
			want:  77,
		},
		{
			input: "99999",
			want:  99,
		},
	}

	for _, tc := range tests {
		result := findCalibrationValue(tc.input)
		if result != tc.want {
			t.Fatalf(`findCalibrationValue(%v) = %v, want %v`, tc.input, result, tc.want)
		}
	}
}

func TestPartOne(t *testing.T) {
	want := 142
	result := partOne(ExampleInput)
	if result != want {
		t.Fatalf(`partOne() = %v, want %v`, result, want)
	}
}
