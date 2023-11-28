package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*TODO: context with deadline*/

// Воркер, завершающий работу при отменет контекста
func CtxWorker(ctx context.Context, channel chan string) {
	for {
		select {
		case <-ctx.Done(): // Нет контекста - нет работы
			fmt.Println("[CONTEXT_WORKER]: exiting...")
			return
		case data := <-channel: // Получаем данные из main()
			fmt.Printf("[CONTEXT_WORKER]: recieved %s \n", data)
		}
	}
}

// Воркер, завершающий свою работу при получаении сигнала из main()
func SignalWorker(signal chan bool, channel chan string) {
	for {
		select {
		case <-signal: // Булевский сигнал есть - нет работы
			fmt.Println("[SIGNAL_WORKER]: exiting...")
			return
		case data := <-channel: // Получаем данные из main()
			fmt.Printf("[SIGNAL_WORKER]: recieved %s \n", data)
		}
	}
}

// Воркер, естетсвнным путём завершается, но сообщает об этом в main()
func WGWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[WAIT_GROUP_WORKER]: called")
}

// Воркер, завершающий себя сам после предопределённой работы
func ExitWorker() {
	fmt.Println("[EXIT_WORKER]: called")
	runtime.Goexit()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // Контекст с отменой
	// Возможны и другие отменяемые контексты, принцип будет тем же, но условий выхода будет больше

	ctxChan := make(chan string) // Канал данных из main() в горутину

	go CtxWorker(ctx, ctxChan)

	ctxChan <- "hello to CONTEXT_WORKER"
	cancel()

	signal := make(chan bool)          // Канал сигналов из main() в горутину
	signalChannel := make(chan string) // Канал данных из main() в горутину

	go SignalWorker(signal, signalChannel)

	signalChannel <- "hello to SIGNAL_WORKER"
	signal <- true

	var wg sync.WaitGroup // Инициализируем группу ожидания

	wg.Add(1)

	go WGWorker(&wg)

	wg.Wait()

	go ExitWorker()

	time.Sleep(1 * time.Second)
}
