package utils

func Min[T orderedNumber](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T orderedNumber](x, y T) T {
	if x > y {
		return x
	}
	return y
}

type orderedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}
