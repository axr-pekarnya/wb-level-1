package main

import (
	"fmt"
	"time"
)

// Простейший таймер, считающ время от вызова до выхода из родительской функции
func timer() func() {
	begin := time.Now()

	return func() {
		fmt.Printf("[TIMER]: %v\n", time.Since(begin))
	}
}

func sleep(seconds int) {
	begin := time.Now().Unix()

	for { // Крутимся в цикле, пока разница во времени входа и текущем времени не окажется больше значения
		if time.Now().Unix()-begin > int64(seconds) {
			break
		}
	}
}

func main() {
	defer timer()()

	begin := time.Now()
	sleep(1)
	end := time.Since(begin)

	fmt.Printf("[MAIN]: %v\n", end)
}
