package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	got := EvenOrOdd(10)
	if got != "Even" {
		t.Errorf("expected Even, got %s", got)
	}
}

func TestEvenOrOddOdd(t *testing.T) {
	got := EvenOrOdd(11)
	if got != "Odd" {
		t.Errorf("expected Odd, got %s", got)
	}
}
