package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]int) // Инициализируем

	var wg sync.WaitGroup // Для синхронизации
	var mutex sync.Mutex

	for i := 0; i < 8; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			mutex.Lock() // Блокируем доступ к данным
			data[i] = i  // Безопасно совершаем запись
			data[i-1] = i
			mutex.Unlock() // Разблокируем доступ к данным
		}(i)
	}

	wg.Wait() // Дожидаемся горутину

	for key, value := range data {
		fmt.Printf("%d - %d \n", key, value)
	}
}
