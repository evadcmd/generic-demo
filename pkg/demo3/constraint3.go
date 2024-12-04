package demo3

import "fmt"

// we can not do something like this
// func f(x Constraint) {
// fmt.Print(x.Func1(""))
// }

type ConstraintWithoutTypeSpecified interface {
	Func1(x string) string
	Func2(elems ...float64) float64
}

// It is OK
func F(x ConstraintWithoutTypeSpecified) {
	fmt.Print(x.Func1(""))
}
