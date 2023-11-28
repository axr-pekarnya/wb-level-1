package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Воркер для горутин
func Work(id int, ctx context.Context, c chan string) {
	for {
		select {
		case <-ctx.Done(): // Если контекст отменён, т.е. программа завершилась, выходим
			fmt.Printf("[WORKER | %d]: Exiting...\n", id)
			return
		case msg := <-c: // Получаем сообщение из канала
			fmt.Printf("[WORKER | %d]: %s\n", id, msg)
		}
	}
}

func main() {
	var numOfWorkers int
	fmt.Print("Input number of Workers: ")
	fmt.Scan(&numOfWorkers)

	WorkerChannel := make(chan string)

	ctx, cancel := context.WithCancel(context.Background()) // Контекст с последующей отменой

	for i := 0; i < numOfWorkers; i++ {
		go Work(i, ctx, WorkerChannel) // Пускаем горутины
	}

	signalChannel := make(chan os.Signal, 1)                    // Канал для получаения сигналов от ОС
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM) // Получать оповешение о прерывании в канал выше

	// Отслеживающая горутина
	go func() {
		<-signalChannel // Ловим сигнал о прерывании
		fmt.Println("[MAIN]: Exiting")
		cancel() // Отменяем контекст
		time.Sleep(time.Second)
		os.Exit(0)
	}()

	// Пихаем данные горутинам
	for {
		select {
		case <-ctx.Done():
			return
		default:
			WorkerChannel <- "random string"
			time.Sleep(time.Second)
		}
	}
}
