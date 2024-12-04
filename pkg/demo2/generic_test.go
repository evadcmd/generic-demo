package demo2

import "testing"

func TestSum(t *testing.T) {
	if sum := Sum(1, 2, 3); sum != 6 {
		t.Errorf("wrong answer: %d", sum)
	} else {
		t.Logf("%T", sum)
	}
	if sum := Sum(1.0, 2.0, 3.0); sum != 6.0 {
		t.Errorf("wrong answer: %f", sum)
	} else {
		t.Logf("%T", sum)
	}
	// !
	if sum := Sum(1, 2.0, 3.0); sum != 6.0 {
		t.Errorf("wrong answer: %f", sum)
	} else {
		t.Logf("%T", sum)
	}
}

func TestAdd1(t *testing.T) {
	x := 1 + 1.0
	t.Logf("%T", x)
}

func TestAdd2(t *testing.T) {
	// x := 1
	// y := 1.0
	// not allowed
	// z := x + y
}
