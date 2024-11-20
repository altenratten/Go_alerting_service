package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func whatWeather(city string) (string, error) {
    baseURL := "https://bba4h86i8icpvi5ot97n.containers.yandexcloud.net/go/weather"
    params := url.Values{}
    params.Add("0", "")
    params.Add("M", "")
    params.Add("city", city)

    urlInstance, err := url.Parse(baseURL)
    // обработайте ошибку и возвратите новую ошибку c текстом следующего формата
    // ошибка преобразования URL: <error>
    if err != nil {
        err := fmt.Errorf("ошибка преобразования URL: %w", err)
        return "", err
    }
    urlInstance.RawQuery = params.Encode()

    response, err := http.Get(urlInstance.String())
    // обработайте ошибку и возвратите новую ошибку c текстом следующего формата
    // сетевая ошибка: <error>
    if err != nil{
        err := fmt.Errorf("сетевая ошибка: %v", err)
        return "", err
    }

    body, err := ioutil.ReadAll(response.Body)
  // обработайте ошибку и возвратите новую ошибку c текстом следующего формата
    // ошибка на сервере погоды: <error>
    if err != nil{
        err := fmt.Errorf("ошибка на сервере погоды: %v", err)
        return "", err
    }
    content := string(body)
    // возвратите результат без ошибки
    return content, nil
}

func main() {
    cities := []string{
        "Омск",
        "Калининград",
        "Челябинск",
        "Владивосток",
        "Красноярск",
        "Москва",
        "Екатеринбург",
    }
  
      fmt.Println("Погода в городах:")
    for _, city := range cities {
        cityWeather, err := whatWeather(city)
        // обработайте ошибку
        // если ошибка произошла, то распечатайте название города и текст ошибки
        // и пропустите итерацию цикла
    if err != nil{
		continue
    }
        fmt.Println(city, cityWeather)
    }
}