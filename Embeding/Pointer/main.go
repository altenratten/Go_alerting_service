package main

import "fmt"

// Hero описывает главного героя игры
type Hero struct {
	Name    string // имя героя
	Health  int    // здоровье
	Damage  int    // урон
	Defense int    // защита
}

// ChangeHealth изменяет здоровье героя на указанное значение.
// вставьте метод ChangeHealth
// ...
func (h *Hero) ChangeHealth(diff int) {
	h.Health += diff
}

func main() {
	hero := &Hero{
		Name:   "Илья Муромец",
		Health: 100500,
	}
	for i := -100; i <= 150; i += 20 {
		hero.ChangeHealth(i)
	}
	fmt.Println(hero.Health) // должна вывести 100760
}
