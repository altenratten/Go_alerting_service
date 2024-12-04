package main

import "fmt"

func main() {
    weight := 75.0  // вес
    height := 175.0 // рост
    steps := 8462.0 // количество пройденных за день шагов
    hours := 2.0    // время движения в часах
    lenStep := 0.65 // длина одного шага в метрах
    
    // напишите формулу расчёта пройденных километров
    dist := steps * lenStep / 1000
    
    meanSpeed := dist / hours
    minutes := hours * 60
    
    spentCalories := (0.035 * weight + (meanSpeed * meanSpeed / height) * 0.029 * weight) * minutes
    
    // здесь выведите строку в нужном формате
    // ...
	fmt.Printf("Сегодня вы прошли %.3f км и затратили %.3f килокалорий./n", dist, spentCalories)
}