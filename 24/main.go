package main

import (
	"fmt"
	"math"
)

// Заданная структура
type Point struct {
	x float64
	y float64
}

// Заданный конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Метод, считающий расстояние до некой точки
func (p1 *Point) DistanceTo(p2 *Point) float64 {
	// Считаем разность координат
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY) // Возвращаем значение по формуле расстояние
}

// Функция, считающая расстояние между двумя точками
func DistanceBetween(p1 *Point, p2 *Point) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(1, 1)

	fmt.Println(DistanceBetween(p1, p2), p1.DistanceTo(p2), p2.DistanceTo(p1))
}
