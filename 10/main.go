package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64) // Будем хранить данные в мапе

	for _, temp := range temperatures {
		key := int(math.Round(temp/10) * 10) // Получаем ближайшее число, кратное десяти

		if math.Abs(float64(key)) > math.Abs(temp) { // Если округлились в большую сторону
			if key > 0 { // Поправляем это, учитывая знак
				key -= 10
			} else {
				key += 10
			}

		}

		groups[key] = append(groups[key], temp)
	}

	for key, value := range groups {
		fmt.Println(key, value)
	}
}
