package main

import "fmt"

// Производим слияние элементов до i, не включая, и оставшихся элементов после i, не включая
// [a, i) U (i, b]

func Pop(data []int, i int) {
	data = append(data[:i], data[i+1:]...)
}

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	Pop(data, 3)
	fmt.Println(data)
}
