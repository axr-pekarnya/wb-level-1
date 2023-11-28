package main

import "fmt"

func main() {
	a := 5
	b := 3

	fmt.Println(a, b)

	// Арифметический метод со сложением
	a += b
	b = a - b
	a -= b

	fmt.Println(a, b)

	// Арифметический метод с делением
	a *= b
	b = a / b
	a /= b

	fmt.Println(a, b)

	// Логический метод с исключающим ИЛИ (XOR)
	a ^= b
	b ^= a
	a ^= b

	fmt.Println(a, b)

	a, b = b, a // Самый просто метод, хотя на стеке должны быть временные переменные

	fmt.Println(a, b)
}
