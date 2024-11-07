package main

import (
	"fmt"
	"sort"
)

// создайте функцию getSortedKeys() ниже
// эта функция должна принимать мапу, а возвращать отсортированный слайс ключей этой мапы
// она нужна, чтобы можно было вывести ключи мапы в отсортированном порядке
// так как мапа сама по себе всегда неотсортирована
func getSortedKeys (years map[int]string) []int{
    keys := make([]int, 0)

    for year := range years {
        keys = append(keys, year)
    }
    sort.Ints(keys)
    return keys
}

func main() {
    moves := map[int]string{
		1991: "Терминатор 2: Судный день",
		1984: "Терминатор",
		2009: "Терминатор: Да придет спаситель",
		2003: "Терминатор 3: Восстание машин",
		2015: "Терминатор: Генезис",
		2019: "Терминатор: Темные судьбы",
	}

	sortedMoves := getSortedKeys(moves)

	fmt.Println("Как смотреть франшизу Терминатор")

	fmt.Println()

	for _, year := range sortedMoves {
		fmt.Println("Год:", year, "Фильм:", moves[year])
	}
}