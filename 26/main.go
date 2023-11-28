package main

import (
	"fmt"
	"strings"
)

func isSet(data string) bool {
	mem := make(map[rune]bool) // Тут будем хранить уже найденные символы строки

	data = strings.ToLower(data) // Преобразуем все символы строки в нижний регистр для регистронезависимости

	for _, char := range data {
		if _, status := mem[char]; status {
			return false // Если в мапе уже есть символ, возращаем значение false
		}

		mem[char] = true // Запоминаем символ
	}

	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "aA"

	fmt.Println(isSet(s1), isSet(s2), isSet(s3), isSet(s4))
}
