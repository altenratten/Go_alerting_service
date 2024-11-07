package main

import (
	"fmt"
	"time"
)

var database = map[string]string{
	"Серёга": "Омск",
	"Соня":   "Москва",
	"Дима":   "Челябинск",
	"Алина":  "Красноярск",
	"Егор":   "Пермь",
}

var offsetUTC = map[string]int{
	"Санкт-Петербург": 3,
	"Москва":          3,
	"Самара":          4,
	"Новосибирск":     7,
	"Екатеринбург":    5,
	"Нижний Новгород": 3,
	"Казань":          3,
	"Челябинск":       5,
	"Омск":            6,
	"Ростов-на-Дону":  3,
	"Уфа":             5,
	"Красноярск":      7,
	"Пермь":           5,
	"Воронеж":         3,
	"Волгоград":       3,
	"Краснодар":       3,
	"Калининград":     2,
}

func whatTime(friend string) string {
	utcTime := time.Now().UTC()
	city, found := database[friend]
	if !found {
		return fmt.Sprintf("У тебя нет друга по имени %s", friend)
	}
	// вычислите время в городе друга и возвратите его в виде строки,
	// используя метод `Format()`
	// ...
	cityTime := utcTime.Add(time.Duration(offsetUTC[city]) * time.Hour)
	return  cityTime.Format("15:04")
}

func main() {
	fmt.Println(whatTime("Соня"))
}
