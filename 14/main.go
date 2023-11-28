package main

import "fmt"

func typeOf(tmp interface{}) string {
	switch tmp.(type) { // Получаем тип из interface{}
	case int:
		return "int"
	case bool:
		return "bool"
	case float64:
		return "float64"
	case string:
		return "string"
	case chan struct{}:
		return "chan"
	default:
		return "uknown type"
	}
}

func main() {
	var tmp interface{}

	tmp = 1
	fmt.Println(typeOf(tmp))

	tmp = true
	fmt.Println(typeOf(tmp))

	tmp = 0.1
	fmt.Println(typeOf(tmp))

	tmp = "kek"
	fmt.Println(typeOf(tmp))

	tmp = make(chan struct{})
	fmt.Println(typeOf(tmp))
}
