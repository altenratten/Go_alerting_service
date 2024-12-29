package main

import (
	"fmt"
)

type Hero struct {
	Name     string
	Health   int
	Location string
}

// ниже объявите метод
func (h *Hero) MoveTo(Location string) {
	h.Location = Location
	fmt.Println(h.Name, "переместился в", Location)
}

func main() {
	myHero := Hero{Name: "Артур", Health: 100}
	// вызовите метод для myHero
	myHero.MoveTo("тронный зал")
}
