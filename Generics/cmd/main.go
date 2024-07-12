package main

import (
	"Generics/internal/service"
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println()
	a := 1
	b := 2
	af := 1.1
	bf := 1.2
	a64 := 3
	b64 := 4

	//problen
	intSum := service.AddInt(a, b)
	floatSum := service.AddFloat(af, bf)

	fmt.Printf("(AddInt) Sum of int -> %d\n", intSum)
	fmt.Printf("(AddFloat) Sum of floats -> %.2f\n", floatSum)

	//solutions
	intSum = service.AddIntAndFloat[int](a, b)
	floatSum = service.AddIntAndFloat[float64](af, bf)

	fmt.Printf("(AddIntAndFloat) Sum of int -> %d\n", intSum)
	fmt.Printf("(AddIntAndFloat) Sum of floats -> %.2f\n", floatSum)

	intSum = service.AddAllIntAndFloat(a, b)
	floatSum = service.AddAllIntAndFloat(af, bf)
	intSum2 := service.AddAllIntAndFloat(a64, b64)

	fmt.Printf("(AddAllIntAndFloat) Sum of int -> %d\n", intSum)
	fmt.Printf("(AddAllIntAndFloat) Sum of float -> %f\n", floatSum)
	fmt.Printf("(AddAllIntAndFloat) Sum of int64 -> %d\n", intSum2)

	intSum = service.AddAllIntAndFloatInterface(a, b)
	floatSum = service.AddAllIntAndFloatInterface(af, bf)
	intSum2 = service.AddAllIntAndFloatInterface(a64, b64)

	fmt.Printf("(AddAllIntAndFloatInterface) Sum of int -> %d\n", intSum)
	fmt.Printf("(AddAllIntAndFloatInterface) Sum of float -> %f\n", floatSum)
	fmt.Printf("(AddAllIntAndFloatInterface) Sum of int64 -> %d\n", intSum2)

	userId1 := service.UserId(1)
	userId2 := service.UserId(2)

	userIdSum := service.AddAllIntAndFloatMatch[service.UserId](userId1, userId2)
	fmt.Printf("(AddAllIntAndFloatMatch) Sum of user ids -> %d\n", userIdSum)

	//--------------------------------------------------------------------------------//
	//problem
	//input [1 ,2, 3], (n) = n * 2
	//output [2, 4, 6]
	fmt.Println("--------------------------------------------------------------")
	intArray := []int{1, 2, 3}
	resultArray := service.MapInts(intArray, func(n int) int {
		return n * 2
	})
	fmt.Printf("Input %d\n", intArray)
	fmt.Printf("(MapValues) %d\n", resultArray)

	resultAnyArray := service.MapAny([]float64{1.1, 1.2, 1.3}, func(f float64) float64 {
		return f * 2
	})

	fmt.Printf("Input %.2f\n", []float64{1.1, 1.2, 1.2})
	fmt.Printf("(MapAny) %.2f\n", resultAnyArray)

	//--------------------------------------------------------------------------------//
	fmt.Println("--------------------------------------------------------------")
	calStruct1 := service.CalStruct{
		Id:   1,
		Name: "test2",
		Data: "hello",
	}
	fmt.Printf("calStruct1 %v\n", calStruct1)
	calStruct2 := service.CalStruct{
		Id:   2,
		Name: "test2",
		Data: 2,
	}
	fmt.Printf("calStruct2 %v\n", calStruct2)

	calStruct3 := service.CalStruct2[string]{
		Id:   3,
		Name: "test3",
		Data: "hello",
	}
	fmt.Printf("calStruct3 %v\n", calStruct3)
	calStruct4 := service.CalStruct2[int]{
		Id:   4,
		Name: "test4",
		Data: 4,
	}
	fmt.Printf("calStruct4 %v\n", calStruct4)

	//--------------------------------------------------------------------------------//
	fmt.Println("--------------------------------------------------------------")
	cusMap1 := make(service.CustomMap[int, string])
	cusMap1[1] = "3"
	fmt.Printf("Custom map 1 %v\n", cusMap1)

	cusMap2 := make(service.CustomMap[string, int])
	cusMap2["1"] = 3
	fmt.Printf("Custom map 2 %v\n", cusMap2)

	cusMap3 := make(service.CustomMap[float32, string])
	cusMap3[1.2] = "test"
	fmt.Printf("Custom map 3 %v\n", cusMap3)

}
