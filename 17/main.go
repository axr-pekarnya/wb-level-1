package main

import "fmt"

func BinarySearch(key int, data []int) int {

	left := 0
	right := len(data) - 1

	for left <= right {
		median := (left + right) / 2 // Ищем середину

		if data[median] < key { // Расмматриваем левую половину, если значение там
			left = median + 1
		} else {
			right = median - 1 // Иначе рассматриваем правую половину
		}
	}

	if left == len(data) || data[left] != key { // Если ключ не найден
		return -1
	}

	return left // В нашей индексации при успешном нахождении индекс ключа лежит в левой границе
}

func main() {
	data := []int{1, 3, 5, 7, 9, 12, 14, 16, 18}

	for _, elem := range data {
		fmt.Printf("%d ", BinarySearch(elem, data))
	}

	fmt.Println()
}
