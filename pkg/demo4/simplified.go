package demo4

func f1[Items interface{ ~[]Item }, Item interface{}](items Items) {

}

func f2[Items ~[]Item, Item interface{}](items Items) {

}

func f3[Items ~[]Item, Item any](items Items) {

}

// not the same
func f4[Items ~[]any](items Items) {

}
