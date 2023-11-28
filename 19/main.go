package main

import "fmt"

func reverse(data string) string {
	runes := []rune(data) // Используем руны для символов юникода
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 { // Здесь два счётчика - с начала слайса и с конца
		runes[i], runes[j] = runes[j], runes[i] // Меняем местами i-ю руну с начала и j-ю с конца
	}

	return string(runes)
}

func main() {
	data := "Ну и дичь 🚼🚼🚼"

	revData := reverse(data)

	fmt.Println(revData)
}
