package main

var justString string

// Поскольку мы взяли не всю строку v, может произойти утечка памяти

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100]) // Создадим новую строку, не ссылающуюся на v

}

// Второй вариант - вернуть срез из строки

func someFunc1() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func main() {
	someFunc()
	justString := someFunc1()
}
