func generic

```
func F[T C](arg1 T) (res T, error)
```

struct generic

```
func F[T C](arg1 T) (res T, error)
```

what is C?

generally speaking, C can be in the following format

```
// (A OR B OR C) AND D AND E
type ComplicatedConstraint[T Constraint] interface {
	~float32 | ~float64 // A OR B OR C
	f1(v T) T  // D
	f2(v T) T  // E
}
```

~float32 means types that are alias of float32

for example

```
type Float float32
```
