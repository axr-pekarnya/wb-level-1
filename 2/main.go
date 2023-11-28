package main

import (
	"fmt"
	"os"
	"sync"
)

// Вариант 1: Простая запись в канал
func Calculate1(num int, c chan int) {
	c <- num * num
}

// Структура для асинхронной записи
type AsyncWriter struct {
	Mutex sync.Mutex
}

// Конструктор
func NewAsyncWriter() *AsyncWriter {
	return &AsyncWriter{Mutex: sync.Mutex{}}
}

// Запись с удержанием горутин
func (aw *AsyncWriter) Write(data int) {
	aw.Mutex.Lock()
	fmt.Fprintf(os.Stdout, "%d ", data)
	aw.Mutex.Unlock()
}

var AsWriter *AsyncWriter = NewAsyncWriter()

// Безопасно пишем в stdout, WaitGroup для ожидания завершения работы горутин
func Calculate2(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	AsWriter.Write(num * num)
}

//ToDo bufio with mutex as argument

func main() {
	data := []int{2, 4, 6, 8, 10}
	c := make(chan int)

	for _, num := range data {
		go Calculate1(num, c)
	}

	for i := 0; i < len(data); i++ {
		fmt.Fprintf(os.Stdout, "%d ", <-c)
	}

	fmt.Fprintln(os.Stdout)

	var wg sync.WaitGroup

	for _, num := range data {
		wg.Add(1)
		go Calculate2(&wg, num)
	}

	wg.Wait()

	fmt.Fprintln(os.Stdout)
}
