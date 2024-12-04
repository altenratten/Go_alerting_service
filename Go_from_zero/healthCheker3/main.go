package main

import "fmt"

func main() {
    weight := 75.0  // вес
    height := 175.0 // рост
    steps := 8462.0 // количество пройденных за день шагов
    hours := 2.0    // время движения в часах
    lenStep := 0.65 // длина одного шага в метрах
    
    dist := steps * lenStep / 1000

    var achievement string
    switch (true) {
	case dist > 6.5:
		achievement = "Отличный результат! Цель достигнута."
	case dist > 3.9:
		achievement = "Неплохо! День был продуктивным."
	case dist > 2.0:
		achievement = "Маловато, но завтра наверстаем!"
	default:
		achievement = "Лежать тоже полезно. Главное — участие, а не победа!"
	}

    meanSpeed := dist / hours
    minutes := hours * 60
    
    spentCalories := (0.035 * weight + (meanSpeed * meanSpeed / height) *
        0.029 * weight) * minutes

    // добавьте вывод результата, можно использовать строку
    // с обратными кавычками
    // ...
	fmt.Printf("Сегодня вы прошли %.0f.\nПройденная дистанция %.2f км.\nВы сожгли %.2f ккал. килокалорий.\n%s\n", steps, dist, spentCalories,achievement)
}
