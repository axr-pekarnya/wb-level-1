package main

import "fmt"

func SetBit(num int64, bitNum int, bitValue int) int64 {
	switch bitValue {
	case 1:
		return num | (1 << bitNum) // Применяем ИЛИ в поизиции bitNum
	case 0:
		return num &^ (1 << bitNum) // Применяем исключающее ИЛИ (XOR) в позиции bitNum
	default:
		return 0
	}
}

func main() {
	var num int64 = 257

	fmt.Printf("%d \n", SetBit(num, 8, 0))
}
