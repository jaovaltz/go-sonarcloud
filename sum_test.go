package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(1, 2)

	if result != 3 {
		t.Error("Expected 3, got ", result)
	}
}

func TestSub(t *testing.T) {
	
	result := sub(1, 2)

	if result != -1 {
		t.Error("Expected -1, got ", result)
	}
}

func TestMul(t *testing.T) {
	
	result := mul(2, 2)

	if result != 4 {
		t.Error("Expected 2, got ", result)
	}
}

func TestDiv(t *testing.T) {
	
	result := div(2, 2)

	if result != 1 {
		t.Error("Expected 1, got ", result)
	}
}