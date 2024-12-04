package demo3

type Item struct {
	name  string
	price float64
}

type Container interface {
	~[]Item
	total() float64
}

type Box []Item

func (b Box) total() (t float64) {
	for _, item := range b {
		t += item.price
	}
	return
}

func sum[T Container](containers ...T) (sum float64) {
	for _, c := range containers {
		sum += c.total()
	}
	return
}

// invalid receiver type []Item
/*
func (items []Item) total() (t float64) {
	for _, item := range items {
		t += item.price
	}
	return
}
*/
