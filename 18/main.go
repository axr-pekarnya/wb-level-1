package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Структура счётчика
type Counter struct {
	data  int
	mutex sync.Mutex
}

// Конструктор счётчика
func NewCounter() *Counter {
	return &Counter{
		data:  0, // Изначально значение равно нулю
		mutex: sync.Mutex{},
	}
}

// Метод инкрементированяи счётчика
func (c *Counter) Add() {
	c.mutex.Lock()   // Блокируемся
	c.data++         // Инкрементируем
	c.mutex.Unlock() // Пускаем другие горутины
}

func main() {
	cnt := NewCounter()
	badCnt := 0                // ЭТО НЕ СЧЁТЧИК, это мальчик для битья - его значение неверно
	var atomicCnt atomic.Int64 // Встроенный вариант - использование атомиков

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			cnt.Add()
			badCnt++
			atomicCnt.Add(1)
		}()
	}

	wg.Wait()

	fmt.Println(cnt.data, badCnt, atomicCnt.Load())
}
