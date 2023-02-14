package calcy

import "testing"

func TestAddition(t *testing.T) {
	result := Addition{}.Calculate(5, 4)
	if result != 9 {
		t.Error("want 9, got", result)
	}
}
func TestSubtraction(t *testing.T) {
	result := Subtraction{}.Calculate(5, 4)
	if result != 1 {
		t.Error("want 1, got", result)
	}
}
func TestMultiplication(t *testing.T) {
	result := Multiplication{}.Calculate(4, 5)
	if result != 20 {
		t.Error("want 20, got", result)
	}
}
func TestDivision(t *testing.T) {
	result := Division{}.Calculate(4, 5)
	if result != 0 {
		t.Error("want 0, got", result)
	}
}
func TestBogus(t *testing.T) {
	result := Bogus{Offset: 42}.Calculate(4, 5)
	if result != 42+9 {
		t.Error("want 51, got", result)
	}
}
