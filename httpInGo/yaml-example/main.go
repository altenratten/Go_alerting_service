package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3" // импортируем пакет для работы с YAML
)

// Artist содержит данные об артисте
type Artist struct {
	ID    int      `yaml:"id"`
	Name  string   `yaml:"name"`
	Genre string   `yaml:"genre"`
	Songs []string `yaml:"songs"`
}

func main() {
	// создаём переменную типа Artist, чтобы конвертировать эти данные в YAML
	artist := Artist{
		ID:    1,
		Name:  "30 seconds to Mars",
		Genre: "rock",
		Songs: []string{
			`The Kill`,
			`A Beautiful Lie`,
			`Attack`,
			`Live Like A Dream`,
		},
	}

	// сериализуем данные из переменной artist в слайс байт
	// & означает, что мы передаём указатель на artist
	yamlData, err := yaml.Marshal(&artist)
	if err != nil {
		fmt.Printf("ошибка при сериализации в yaml: %s", err.Error())
		return
	}

	// выводим результат в виде слайса байт
	fmt.Println(yamlData)
	// и в виде строки
	fmt.Println(string(yamlData))

	// создаём yaml-файл artist.yaml
	f, err := os.Create("artist.yaml")
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return
	}
	// когда программа завершится, надо закрыть дескриптор файла
	defer f.Close()

	// записываем слайс байт в файл
	_, err = f.Write(yamlData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return
	}
}
