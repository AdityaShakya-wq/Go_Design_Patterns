package service

type CustomType interface {
	int | int32 | int64 | float32 | float64 | []byte | []rune | string
}
type CalStruct struct {
	Id   int
	Name string
	Data interface{}
}

type CalStruct2[T CustomType] struct {
	Id   int
	Name string
	Data T
}
