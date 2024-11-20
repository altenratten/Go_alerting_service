package main

import (
    "fmt"
    "time"
)

func whatTime(city string) time.Time {
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
    // напишите код, который берёт текущее UTC-время
    // и добавляет к нему сдвиг города в часах
    // ...
	utc := time.Now().UTC()
	cityTime := utc.Add( time.Duration(offsetUTC[city]) * time.Hour )
	// cityTime := 
	return cityTime
}

func main() {
    t := whatTime("Екатеринбург")
    fmt.Printf("В Екатеринбурге %d ч.\n", t.Hour())
}