package demo3

type Constraint interface {
	~int | ~int32 | ~int64
	Func1(x string) string
	Func2(elems ...float64) float64
}

type ComplicatedConstraint[T Constraint] interface {
	~float32 | ~float64
	Transform(v T) T
}

func X[T Constraint](ts ...T) {

}

type Int int

func (Int) Func1(x string) string {
	return x
}

func (Int) Func2(fs ...float64) (sum float64) {
	for _, v := range fs {
		sum += v
	}
	return
}

func Y[T ComplicatedConstraint[Int]](ts ...T) {

}
