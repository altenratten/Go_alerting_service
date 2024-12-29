package main

import (
 "fmt"
 "math"
)

var Message = `Добро пожаловать в самую лучшую программу для вычисления
квадратного корня из заданного числа`


func calculateSquareRoot(number float64) float64 {
    return math.Sqrt(number)
}

// Вычисление корня
func calc(number float64) {
    if number < 0 {
        return
    } 
	fmt.Println("Мы вычислили квадратный корень из введённого вами числа. Это будет:", calculateSquareRoot(number))
}

func main() {
    fmt.Println(Message)
    calc(25.5)
}