package main

import "fmt"

func main() {
	// Используем в качестве значения struct{}, потому что нам и не нужны значения, можно и bool
	set1 := map[int]struct{}{
		2:  {},
		4:  {},
		6:  {},
		8:  {},
		10: {},
		12: {},
	}

	set2 := map[int]struct{}{
		3:  {},
		6:  {},
		9:  {},
		12: {},
	}

	result := make(map[int]struct{})

	for key, value := range set1 { // Проходимся по одному множеству
		_, exists := set2[key] // Ищем элемент из первого множества во втором

		if exists {
			result[key] = value // Если элемент по значению в обоих множествах, добавляем в результат.
		}
	}

	for key, value := range result {
		fmt.Println(key, value)
	}
}
