package main

import "fmt"

type Thing struct {
	Name   string
	Weight int
}

type Room struct {
	Name   string
	Things []Thing
}

type House struct {
	Name  string
	Rooms [][]Room
}

func main() {
	house := House{
		Name: "Дом v1",
		// инициализируйте house нужными значениями
		// ...
		Rooms: [][]Room{
			{ // Подвал
				{
					Name: "Кладовка", Things: []Thing{
						{Name: "топор", Weight: 3000},
						{Name: "фонарик"},
						{Name: "брелок"},
					},
				},
				{
					Name: "Котельная", Things: []Thing{
						{Name: "верёвка", Weight: 200},
						{Name: "рюкзак", Weight: 500},
					},
				},
			},
			{ // 1 этаж
				{
					Name: "Столовая", Things: []Thing{
						{Name: "карандаш"},
						{Name: "кольцо"},
						{Name: "карта"},
						{Name: "бинт"},
					},
				},
			},
		},
	}
	fmt.Println(house)
}
