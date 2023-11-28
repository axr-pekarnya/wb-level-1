package main

import "fmt"

// Пишем в канал квадрат числа
func CalcSquare(num int, c chan int) {
	c <- num * num
}

func main() {
	data := []int{2, 4, 6, 8, 10}
	c := make(chan int)
	result := 0

	for _, num := range data {
		go CalcSquare(num, c)
	}

	for i := 0; i < len(data); i++ {
		result += <-c // Получаем квадраты от горутин
	}

	fmt.Println(result)
}
