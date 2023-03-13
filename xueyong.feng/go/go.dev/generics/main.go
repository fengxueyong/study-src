package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, value := range m {
		sum += value
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, value := range m {
		sum += value
	}
	return sum
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	mapInts := map[string]int64{
		"a": 100,
		"b": 1000,
	}

	mapFloats := map[string]float64{
		"a": 35.98,
		"b": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(mapInts),
		SumFloats(mapFloats))

	// 这里在调用的时候显式明确泛型的具体类型
	// 下一条语句你将看到省略的情况
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](mapInts),
		SumIntsOrFloats[string, float64](mapFloats))

	// 和上面同样的输出结果。
	// 因为函数有参数，可以通过参数推导出具体的类型，所以可以省略显式的具体类型
	// 但是如果函数没有参数，那就无法推导出，那么就不能省略了。
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(mapInts),
		SumIntsOrFloats(mapFloats))

	// 使用接口定义
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(mapInts),
		SumNumbers(mapFloats))
}
