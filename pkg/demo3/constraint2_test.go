package demo3

import "testing"

func TestContainerSum(t *testing.T) {
	box := Box([]Item{{name: "banana", price: 2.0}, {name: "orange", price: 2.5}})
	s := sum(box)
	if s != 4.5 {
		t.Errorf("wrong answer: %f", s)
	}
}
