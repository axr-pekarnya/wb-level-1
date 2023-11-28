package main

import (
	"fmt"
	"math/rand"
)

// Разбиваем срез data на две части относительно выбранного опорного элемента (pivot)
func partition(data []int, left, right int) ([]int, int) {
	pivot := data[right] // Выбираем опорный элемент
	i := left

	for j := left; j < right; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i] // Переставляем элементы меньше опорного элемента
			i++
		}
	}

	data[i], data[right] = data[right], data[i] // Переставляем опорный элемент на его правильное место

	return data, i
}

// Сортируем в указанном диапазоне
func quickSort(data []int, left, right int) []int {
	if left < right {
		var p int
		data, p = partition(data, left, right)
		data = quickSort(data, left, p-1)  // Сортируем левую часть
		data = quickSort(data, p+1, right) // Сортируем правую часть
	}
	return data
}

// QuickSort - публичная функция для сортировки, запускающая рекурсию
func QuickSort(data []int) []int {
	return quickSort(data, 0, len(data)-1)
}

func main() {
	var data []int

	for i := 0; i < 10; i++ {
		data = append(data, rand.Intn(20)) // Заполняем случайными числами
		fmt.Printf("%d ", data[i])
	}

	fmt.Println()

	result := QuickSort(data)

	for _, elem := range result {
		fmt.Printf("%d ", elem)
	}

	fmt.Println()
}
