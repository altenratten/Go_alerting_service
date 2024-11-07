package main

import (
	"fmt"
	"sort"
	"strings"
)

var database = map[string]string{
	"Серёга": "Омск",
	"Соня":   "Москва",
	"Миша":   "Оренбург",
	"Дима":   "Челябинск",
	"Алина":  "Красноярск",
	"Егор":   "Пермь",
	"Коля":   "Екатеринбург",
}

// новая функция, она возвращает правильное словосочетание, склоняя слово «друзья»
// в зависимости от того, какое число передано в аргументе friendsCount
func formatFriendsCount(friendsCount int) string {
	if friendsCount == 1 {
		return "1 друг"
	}
	if friendsCount >= 2 && friendsCount <= 4 {
		return fmt.Sprintf("%d друга", friendsCount)
	}
	return fmt.Sprintf("%d друзей", friendsCount)
}

func processAlice(query string) string {
	if query == "сколько у меня друзей?" {
		// вызовите функцию formatFriendsCount() и передайте в неё friendsCount;
		// отредактируйте строку ниже: в ней должно использоваться выражение,
		// которое вернёт функция formatFriendsCount()
		friendsCount := formatFriendsCount(len(database)) // friendsCount должна содержать общее количество друзей в database
		return fmt.Sprintf("У тебя %s.", friendsCount)
	}
	if query == "кто все мои друзья?" {
		var friends []string
		for name := range database {
			friends = append(friends, name)
		}
		sort.Strings(friends)
		return fmt.Sprintf("Твои друзья: %s", strings.Join(friends, ", "))
	}
	if query == "где все мои друзья?" {
		uniqueCities := map[string]int{}
		for _, city := range database {
			// заполняем мапу без дублирования городов
			uniqueCities[city] = 1
		}
		var cities []string
		// получаем уникальные названия городов
		for city := range uniqueCities {
			cities = append(cities, city)
		}
		sort.Strings(cities)
		return fmt.Sprintf("Твои друзья в городах: %s", strings.Join(cities, ", "))
	}
	return "неизвестный запрос"
}

func processQuery(query string) string {
	queryAndName := strings.Split(query, ", ")
	name := queryAndName[0]
	queryOnly := queryAndName[1]
	if name == "Алиса"{
		return processAlice(queryOnly)
	} else {
		return processFriend(name, queryOnly)
	}
	return ""
}

func processFriend(name string, query string) string {
	city, include := database[name]
	if include && query == "ты где?" {
		return fmt.Sprintf("%s в городе: %s", name, city)
	} else if include && query != "ты где?" {
		return "неизвестный запрос"
	}

	return fmt.Sprintf("У тебя нет друга по имени %s", name)
}

func main() {
	fmt.Println("Привет, я Алиса!")
	fmt.Println(processQuery("Алиса, сколько у меня друзей?"))
	fmt.Println(processQuery("Алиса, кто все мои друзья?"))
	fmt.Println(processQuery("Алиса, где все мои друзья?"))
	fmt.Println(processQuery("Алиса, кто виноват?"))
	fmt.Println(processQuery("Соня, ты где?"))
	fmt.Println(processQuery("Коля, что делать?"))
	fmt.Println(processQuery("Антон, ты где?")) 
}
