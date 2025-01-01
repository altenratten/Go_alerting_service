package main

import (
	"fmt"
)

// Hero описывает героя с общими полями для всех классов
// в этой структуре ничего не поменялось
type Hero struct {
	Health int // здоровье
	Damage int // урон
	Def    int // защита
}

// в следующие три структуры добавим урон `Damage`

// Warrior описывает класс `Воин`
type Warrior struct {
	Hero             // встроенная структура Hero
	Name      string // имя
	ClassName string // имя класса
	Damage    int    // урон для класса Воин
}

// Mage описывает класс `Маг`
type Mage struct {
	Hero             // встроенная структура Hero
	Name      string // имя
	ClassName string // имя класса
	Damage    int    // урон для класса Маг
}

// Healer описывает класс `Лекарь`
type Healer struct {
	Hero             // встроенная структура Hero
	Name      string // имя
	ClassName string // имя класса
	Damage    int    // урон для класса Лекарь
}

func main() {
	// создаем структуру героя для всех классов
	hero := Hero{Health: 100, Damage: 30, Def: 20}

	// создаем воина
	warrior := Warrior{Hero: hero, Name: "Арагорн", ClassName: "Воин", Damage: 40}
	// урон из Hero. Для вызова поля Damage из встроенной структуры Hero надо эту структуру указать
	fmt.Println("Поле Damage структуры Hero", warrior.Hero.Damage)
	// урон из Warrior. А здесь мы просто указываем поле структуры Warrior
	fmt.Println("Поле Damage структуры Warrior", warrior.Damage)
}
