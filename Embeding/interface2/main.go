package main

import "fmt"

// Sword описывает меч.
type Sword struct {
	Power int
}

// Scroll описывает свиток.
type Scroll struct {
	Magic int
}

// опишите интерфейсный тип Loot
// ...
type Loot interface {
	Apply()
}

// добавьте нужный метод для типов Sword и Scroll
// ...
func (sword Sword) Apply() {
	fmt.Println("Меч", sword.Power)
}

func (scroll Scroll) Apply() {
	fmt.Println("Свиток", scroll.Magic)
}

func main() {
	// loot - это слайс интерфейсного типа Loot. Так как типы Sword и Scroll
	// должны удовлетворять этому интерфейсу, то можно использовать эти структуры
	// как элементы слайсы. Этот слайс создан исключительно для проверки того,
	// правильно ли реализован тип Loot и метод Apply() для структур
	loot := []Loot{Sword{Power: 50}, Scroll{Magic: 20}, Scroll{Magic: 70}}

	for _, v := range loot {
		v.Apply()
	}
}
