package mylibrary_test

import (
	"testing"

	math "github.com/dineshnatukula/mathlib/math"
)

func TestAdd(t *testing.T) {
	result := math.Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, 5)
	}

	result = math.Add(2, 8)
	if result != 10 {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, 5)
	}

	result = math.Add(2, -3)
	if result != -1 {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, 5)
	}
}

func TestSub(t *testing.T) {
	result := math.Sub(2, 3)
	if result != -1 {
		t.Errorf("Sub(2, 3) returned %d, expected %d", result, 5)
	}

	result = math.Sub(2, 8)
	if result != -6 {
		t.Errorf("Sub(2, 3) returned %d, expected %d", result, 5)
	}

	result = math.Sub(2, -3)
	if result != 5 {
		t.Errorf("Sub(2, 3) returned %d, expected %d", result, 5)
	}
}
