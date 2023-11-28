package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
)

/*TODO FOR MONEY: длинная арифметика и быстрое преобразование Фурье*/

// Из арифметического выражения вида "a + b" извлекаем операнды и оператор
func ParseExpression(expr string) (*big.Int, *big.Int, string, error) {
	words := strings.Fields(expr)

	op1 := new(big.Int)
	op2 := new(big.Int)

	if len(words) != 3 {
		return op1, op2, "", errors.New("Invalid expression")
	}

	op1.SetString(words[0], 10)
	op2.SetString(words[2], 10)

	return op1, op2, words[1], nil
}

func Operate(expr string) *big.Int {
	a, b, operator, err := ParseExpression(expr)

	if err != nil {
		log.Fatalf("%f", err)
	}

	result := new(big.Int) // Заводим ещё одну переменную big.Int для обработки результата

	switch operator { // Выполняем операции в соответствии со значением оператора строкового типа
	case "+":
		return result.Add(a, b)
	case "-":
		return result.Add(a, b.Neg(b))
	case "*":
		return result.Mul(a, b)
	case "/":
		return result.Div(a, b)
	}

	return result
}

func main() {
	hugeSum := "32000000000000000000 + 64000000000000000000000"
	hugeDif := "64000000000000000000000 - 32000000000000000000000"
	hugeMul := "64000000000000000000000 * 64000000000000000000000"
	hugeDiv := "64000000000000000000000 / 32000000000000000000000"

	sum := Operate(hugeSum)
	dif := Operate(hugeDif)
	mul := Operate(hugeMul)
	div := Operate(hugeDiv)

	fmt.Println(sum, dif, mul, div)
}
