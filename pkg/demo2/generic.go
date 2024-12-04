package demo2

func Sum[I ~int | ~int64 | ~int32 | float32 | float64](elems ...I) (s I) {
	for _, v := range elems {
		s += v
	}
	return
}
