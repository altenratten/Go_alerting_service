package main

import (
    "fmt"
    "strings"
)


func checkLang(language string) string{
	language = strings.ToLower(language)
	if language == "python" {
		fmt.Println(language, " — интерпретируемый!")
	} else {
		fmt.Println(language, " — интерпретируемый!")
	}
}

func main() {
    // строка, которую нужно разбить на слайс
    languages := "C++, Go, Rust, Java, Python"

    str := "— это языки программирования!"
    // выводим сообщение, объединив две строки
    fmt.Println(languages, str)

    slice := strings.Split(languages, ", ") // здесь разбейте строку из languages и создайте слайс

    // ниже используйте for для получения каждого элемента слайса и вывода его на экран
    for _, language := range slice {
		checkLang(language)
    }
}