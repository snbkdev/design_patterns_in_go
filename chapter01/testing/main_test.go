package main

import "testing"

func TestSum(t *testing.T) {
	a := 5
	b := 6
	expected := 11

	res := sum(a, b)

	if res != expected {
		t.Errorf("Our sum function does not work, %d + %d is not %d\n", a, b, res)
	}
}

func TestMultiply(t *testing.T) {
	a := 5
	b := 6
	expected := 30

	res := multiply(a, b)
	if res != expected {
		t.Errorf("Our multiply function does not work, %d * %d is not %d\n", a, b, res)
	}
}