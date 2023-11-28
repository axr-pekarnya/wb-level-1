package main

import "fmt"

type Human struct {
	Id int
}

// Конструктор
func NewHuman(id int) *Human {
	return &Human{Id: id}
}

// Демонастрационный метод структуры Human
func (h *Human) Run() {
	fmt.Printf("Human with id = %d doing something...\n", h.Id)
}

type Action struct {
	Id    int
	Human *Human // Экземпляр встраиваемой структуры
}

// Конструктор с инициализацией Human
func NewAction(ActionId, HumanId int) *Action {
	return &Action{Id: ActionId, Human: NewHuman(HumanId)}
}

// Демонстрационный метод структуры Action
func (a *Action) Run() {
	fmt.Printf("Action with id = %d is calling human...\n", a.Id)
	a.Human.Run()
}

func main() {
	a := NewAction(0, 1)
	a.Run()
}
