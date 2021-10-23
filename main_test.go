package main

import "testing"

func TestFunc(t *testing.T) {
	fibValue := fib(5)
	if fibValue != 21 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", fibValue, 21)
	}

	facValue := factorial(5)
	if facValue != 720 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", facValue, 720)
	}
}

func TestFunc2(t *testing.T) {
	tables := []struct {
		input  int64
		output int64
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{7, 5040},
	}

	for _, table := range tables {
		total := factorial(table.input)
		if total != table.output {
			t.Errorf("Factorial of (%d) was incorrect, got: %d, want: %d.", table.input, table.output, total)
		}
	}
}
