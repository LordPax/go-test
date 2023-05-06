package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 != 3")
	}
}

func TestSub(t *testing.T) {
	if Sub(5, 2) != 3 {
		t.Error("5 - 2 != 3")
	}
}

func TestMul(t *testing.T) {
	if Mul(2, 3) != 6 {
		t.Error("2 * 3 != 6")
	}
}

func TestDiv(t *testing.T) {
	if Div(6, 3) != 2 {
		t.Error("6 / 3 != 2")
	}
	if Div(6, 0) != 0 {
		t.Error("6 / 0 != 0")
	}
}

func TestAvg(t *testing.T) {
	number := []int{1, 2, 3}
	if Avg(number) != 2 {
		t.Error("average of 1, 2, 3 is not equal 2")
	}
}
