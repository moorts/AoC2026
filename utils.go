package main

type Summable interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}

func fmap[IN any, OUT any](slice []IN, f (func(IN) OUT)) (result []OUT) {
	n := len(slice)
	result = make([]OUT, n, n)
	for i, element := range slice {
		result[i] = f(element)
	}
	return
}

func sum[IN Summable](slice []IN) (result IN) {
	for _, element := range slice {
		result += element;
	}
	return
}

func reduce[IN any, OUT any](slice []IN, initial OUT, f func(OUT, IN) OUT)) OUT {
	acc := initial;
	for _, element := range slice {
		acc = f(acc, element)
	}
	return acc
}
