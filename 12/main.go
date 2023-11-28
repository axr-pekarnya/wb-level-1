package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool) // Опять же значения не важны, для разнообразия используем bool

	for _, elem := range data {
		set[elem] = true // По свойствам мапы ключ уникален, поэтому просто записываем
		// Фактически перезаписываем, но т.к. значение не меняется, мы практически ничего не делаем
	}

	for key := range set {
		fmt.Printf("%s ", key)
	}

	fmt.Println()
}
