package service

type UserId int

type Num interface {
	int | int32 | int64 | float32 | float64
}

func AddInt(a int, b int) int {
	return a + b
}

func AddFloat(a float64, b float64) float64 {
	return a + b
}

func AddIntAndFloat[T int | float64](a T, b T) T {
	return a + b
}

func AddAllIntAndFloat[T int | int32 | int64 | float32 | float64](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum = sum + num
	}

	return sum
}

func AddAllIntAndFloatInterface[T Num](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum = sum + num
	}

	return sum
}

func AddAllIntAndFloatMatch[T ~int | float32](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum = sum + num
	}

	return sum
}
