package main

import "fmt"

// Property описывает общие свойства предметов.
type Property struct {
	Name  string // название предмета
	Price int
}

// Loot содержит методы управления инвентарём.
type Loot interface {
	Apply()
	Properties() Property
	Sell() int
}

type Scroll struct {
	Property
}

type Sword struct {
	Property
}

// добавьте нужные методы для типов Scroll и Sword
// ...

func (scroll Scroll) Apply() {
	// ...
}

func (scroll Scroll) Properties() Property {
	return scroll.Property
}

func (scroll Scroll) Sell() int {
	return scroll.Price
}

func (sword Sword) Apply() {
	// ...
}

func (sword Sword) Properties() Property {
	return sword.Property
}

func (sword Sword) Sell() int {
	return sword.Price
}

func main() {
	loot := []Loot{
		Scroll{Property{Name: "Свиток знаний", Price: 5}},
		Sword{Property{Name: "Двуручный меч"}},
	}

	for _, v := range loot {
		fmt.Println(v.Properties())
	}
	fmt.Println("Успех!")
}
