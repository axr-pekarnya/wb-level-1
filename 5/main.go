package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Читает из канала
func Read(ctx context.Context, channel chan int) {
	for {
		select {
		case <-ctx.Done(): // Контекста нет - работы нет
			return
		case data := <-channel:
			fmt.Printf("[READER]: recieved %d\n", data)
			time.Sleep(time.Second)
		}
	}
}

// Пишет в канад
func Write(ctx context.Context, channel chan int) {
	for {
		select {
		case <-ctx.Done(): // Контекста нет - работы нет
			return
		default:
			num := rand.Intn(10) // Генерим случайное
			fmt.Printf("[WRITER]: sending %d\n", num)
			time.Sleep(time.Second)
			channel <- num
		}
	}
}

func main() {
	var duration int

	fmt.Print("Input duration: ")
	fmt.Scan(&duration)

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan int)

	go Read(ctx, c)
	go Write(ctx, c)

	time.Sleep(time.Duration(duration) * time.Second) // Простейший вариант - поспать

	cancel()

	// Вариант хитрее - просрачивающийся контекст
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel1() // Отмена либо по времени, либо при выходе из main()

	go Read(ctx1, c)
	go Write(ctx1, c)

	<-ctx1.Done() // Дожидаемся отмены контекста по времени
}
