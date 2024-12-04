package main

import "fmt"

func main() {
    weight := 75.0  // вес
    height := 175.0 // рост
    dist := 9.75    // расстояние в км
    hours := 2.0    // время движения в часах
    k1 := 0.035     // первый коэффициент
    k2 := 0.029     // второй коэффициент

    meanSpeed := dist / hours
    minutes := hours * 60

    spentCalories := ((weight * k1 ) + ((meanSpeed * meanSpeed) / height) * (weight * k2)) *  minutes // напишите формулу расчёта

    fmt.Printf("%.4f", spentCalories)
}