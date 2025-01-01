package main

import "fmt"

func main() {
	var intSlice []int = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var floatSlice []float32 = []float32{1.1, 2.2, 3.1}
	fmt.Println(sumSlice[float32](floatSlice))

	fmt.Println(isEmpty(floatSlice))

}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func sumSlice[T int | float32](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
