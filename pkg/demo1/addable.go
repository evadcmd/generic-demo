package demo1

import "fmt"

type Addable interface {
	Add(Addable) (Addable, error)
}

// 全てのAddableを同じ型をするのはできない
func Sum(elems ...Addable) (Addable, error) {
	var sum Addable = nil
	if len(elems) == 0 {
		return sum, nil
	}
	sum = elems[0]
	for _, v := range elems[1:] {
		var err error
		sum, err = sum.Add(v)
		if err != nil {
			return nil, err
		}
	}
	return sum, nil
}

type Int int

func (i Int) Add(addable Addable) (Addable, error) {
	switch v := addable.(type) {
	case Int:
		return Int(i + v), nil
	// case Int32...
	default:
		return nil, fmt.Errorf("not supported")
	}
}

// type Int32 int32
// ...
