package main

import "fmt"

// Адаптируем котика Cat в человека Human

// Адаптируемая структура
type Cat struct {
	Name string
	Age  int
}

// Конструктор
func NewCat(name string, age int) *Cat {
	return &Cat{Name: name, Age: age}
}

// Демонстрационный метод
func (c *Cat) SayMeow() {
	fmt.Printf("[%s]: MEEEEEEEEOOOOOOOOOW\n", c.Name)
}

// Целевой интерфейс
type Human interface {
	Work()
}

// Структура адаптера
type HumanAdapter struct {
	Cat *Cat
}

// Конструктор
func NewHumanAdapter(cat *Cat) *HumanAdapter {
	return &HumanAdapter{Cat: cat}
}

// Демонстрационный метод
func (h *HumanAdapter) Work() {
	fmt.Println("I'm going to work.")
}

func main() {
	cat := NewCat("Pubich", 13)
	cat.SayMeow()

	humanAdapter := NewHumanAdapter(cat)
	humanAdapter.Work()
	humanAdapter.Cat.SayMeow()
}
