package service

func MapInts(num []int, mapFunc func(int) int) []int {
	var newValues []int
	for _, v := range num {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}

	return newValues
}

func MapAny[T int | int32 | int64 | float32 | float64](num []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range num {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}

	return newValues
}
