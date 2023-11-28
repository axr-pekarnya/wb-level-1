package main

import (
	"fmt"
	"strings"
)

func reverseWords(data string) string {
	words := strings.Fields(data) // Пользуемся встроенной функцией разделения по пробелам

	var result string

	for i := len(words) - 1; i >= 0; i-- { // Итерируемся по словам с конца
		result += fmt.Sprintf("%s", words[i]) // Записываем в результат

		if i != 0 {
			result += " " // Если это не последнее слово в result, вставляем пробел, чтобы не было пробела в начале строки
		}
	}

	return result
}

func main() {
	data := "Ну и 🚼🚼 дичь!"

	revData := reverseWords(data)

	fmt.Println(revData)
}
