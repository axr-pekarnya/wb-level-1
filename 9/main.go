package main

import (
	"fmt"
	"sync"
)

// Первая функция конвейера: получил - возвёл - отправил
func firstPipe(chanIn <-chan int, chanOut chan<- int) { //chanIn только для чтения, chanOut только для записи
	for num := range chanIn { // Читаем канал, пока он открыт
		chanOut <- num * num
	}

	close(chanOut) // Закрываем канал, когда всё отправили
}

// Вторая функция конвейера: получил - напечатал - сообщил о завершении
func secondPipe(wg *sync.WaitGroup, chanIn <-chan int) { // chanIn только для чтения
	defer wg.Done() // При выходе из функции сообщаем о завершении горутины

	for num := range chanIn {
		fmt.Printf("%d\n", num)
	}
}

// Функция отправки данных на конвейер (в первый канал)
func send(data []int, chanOut chan<- int) {
	for _, elem := range data {
		chanOut <- elem
	}

	close(chanOut) // Закрываем, чтобы fitstPipe понял, что все данные получены
}

func main() {
	data := []int{1, 2, 3, 4}

	var wg sync.WaitGroup
	wg.Add(1)

	toFirst := make(chan int)
	toSecond := make(chan int)

	go send(data, toFirst)
	go firstPipe(toFirst, toSecond)
	go secondPipe(&wg, toSecond)

	wg.Wait()
}
